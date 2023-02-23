'use client'
import DogImageScrolling from "@/app/DogImageScroll";
import React, {useCallback, useEffect, useReducer, useState} from "react";
import {Breed} from "@/app/DogDropDown";
import {uid} from "uid";

interface Props {
    dogCount:number,
    class:string,
    filteredBreeds:Breed[],
    direction:string,
}


interface Update{
    UID: string
    newRow?:string[]
}

function reducer(state:{[uid:string]:string[]},update:Update) : {[id:string]:string[]}{
    if(update.newRow){
        let addNewRow : {[id:string]:string[]} = {}
        addNewRow[update.UID] = update.newRow;
        return {...state,... addNewRow}
    }
    let newState = {...state}
    delete newState[update.UID];
    return newState
}

export default function ManyDogs(props:Props){

    const [loading,setLoading] = useState(true)
    const [dogImageRows,updateDogImages] = useReducer(reducer,{})

    let getRandomInt = useCallback((min:number, max:number,seed:number) => {
        min = Math.ceil(min);
        max = Math.floor(max);
        let x = Math.sin(seed) * 10000;
        let random = x - Math.floor(x);
        return Math.floor(random * (max - min + 1) + min); // The maximum is inclusive and the minimum is inclusive
    },[])


    let getDogRow = useCallback((dogImages:string[],className:string,key:string)=>{
        let dogs : JSX.Element[] = [];
        for(let i=0;i<dogImages.length;i++){
            let seed = parseInt(key)+i
            let speedDeviation = getRandomInt(-15000,0,seed)
            let delayDeviation = getRandomInt(-1000,1000,seed+1)
            let xDeviation = getRandomInt(-100,+100,seed+2)
            let yDeviation = getRandomInt(-100,+100,seed+3)
            dogs.push(DogImageScrolling({deviation:{x:xDeviation,y:yDeviation}, key: i,alt: "", class: `${props.class} drop-shadow-xl`, src: dogImages[i], style: {animationDuration:`${30000+speedDeviation}ms`,animationDelay:`${(delayDeviation).toString()}ms`}}))
        }
        return (
            <div className={className} key={key}>
                {dogs}
            </div>
        );
        },[getRandomInt, props.class])


    let getDogs = useCallback(()=>{
        let className = "absolute grid gap-2"
        if(props.direction == "left" || props.direction == "right"){
            className += " grid-rows-6 h-full"
        }
        else {
            className += " grid-cols-6 w-full"
        }
        let dogRows : JSX.Element[] = [];
        for(let dogImageRow in dogImageRows){
            dogRows.push(getDogRow(dogImageRows[dogImageRow],className,dogImageRow))
        }
        return dogRows
    },[dogImageRows, getDogRow, props.direction])

    let content = useCallback(()=>{
        if(loading){
            return (
                <div className="w-full h-full m-auto flex-col mx-auto">
                    <h1 className="w-fit mx-auto font-bold text-3xl">Arf, Fetching!</h1>
                    <img src="/favicon.svg" className="h-72 w-72 mx-auto" alt=""/>
                </div>
            )
        }
        else{
            return(
                <div className="w-full h-full">
                    {getDogs()}
                </div>
            )
        }
    },[getDogs, loading])

    let addDogRow = useCallback((images:string[]) => {
        let id:string = uid()
        updateDogImages({UID:id,newRow:images})
        setTimeout(()=>{
            updateDogImages({UID:id})

        },30000)
    },[])



    const [dogImagesBuffer, setDogImageBuffer] = useState<string[]>([])

    let addNewDogRow = useCallback(()=>{
        if(props.filteredBreeds.length == 0){
            fetch(`api/dog/breeds/image/random/${props.dogCount}`)
                .then(res => res.json())
                .then(json => {
                    setDogImageBuffer(json['message'])
                    setLoading(false)})
        }
        else {
            for(let Breed in props.filteredBreeds){
                let eachBreed = Math.max(Math.floor(props.dogCount / props.filteredBreeds.length),1)
                fetch(`api/dog/breed/${props.filteredBreeds[Breed].value}/images/random/${eachBreed}`)
                    .then(res => res.json())
                    .then(json => {
                        setDogImageBuffer((prevState => prevState.concat(json["message"])))
                        setLoading(false)
                    })
            }
        }
    },[props.dogCount, props.filteredBreeds])

    let [canLoop,setCanLoop] = useState(true)

    let loopAdd = useCallback(()=>{
        if(!canLoop){
            return
        }
        setCanLoop(false)
        setTimeout(()=>{
            addDogRow(dogImagesBuffer)
            setCanLoop(true)
        }, 1000)
        setDogImageBuffer([])
        addNewDogRow()
    },[addDogRow, addNewDogRow, dogImagesBuffer, canLoop])




    useEffect(()=>{
        if(canLoop){
            loopAdd()
        }
    },[loopAdd, canLoop])

    return(
        <div className="fixed w-full h-full -z-10">
            {content()}
        </div>
    )
}