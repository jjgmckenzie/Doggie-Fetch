'use client'
import DogImageScrolling from "@/app/DogImageScroll";
import React, {Dispatch, SetStateAction, useCallback, useEffect, useReducer, useState} from "react";
import {Breed} from "@/app/Breed";
import {uid} from "uid";

interface Props {
    dogCount:number,
    class:string,
    filteredBreeds:Breed[],
    direction:string,
    animSpeed:number,
    imgSize:number,
    loading:boolean
    setLoading:Dispatch<SetStateAction<boolean>>
}

interface ImageRow{
    animSpeed:number,
    imgSize:number,
    image:string,
}

interface Update{
    UID: string
    newRow?:ImageRow[]
}

function reducer(state:{[uid:string]:ImageRow[]},update:Update) : {[id:string]:ImageRow[]}{
    if(update.newRow){
        let addNewRow : {[id:string]:ImageRow[]} = {}
        addNewRow[update.UID] = update.newRow;
        return {...state,... addNewRow}
    }
    let newState = {...state}
    delete newState[update.UID];
    return newState
}

export default function ManyDogs(props:Props){
    const [dogImageRows,updateDogImages] = useReducer(reducer,{})

    let getRandomInt = useCallback((min:number, max:number,seed:number) => {
        min = Math.ceil(min);
        max = Math.floor(max);
        let x = Math.sin(seed) * 10000;
        let random = x - Math.floor(x);
        return Math.floor(random * (max - min + 1) + min); // The maximum is inclusive and the minimum is inclusive
    },[])


    let getDogRow = useCallback((rows:ImageRow[],key:string)=>{
        let style : React.CSSProperties = {}
        let className = ''
        if(rows.length > 0){
            if(props.direction == "left" || props.direction == "right"){
                className = "flex-col h-full"
            }
            else {
                className="flex w-full"
            }
        }
        let dogs : JSX.Element[] = [];
        for(let i=0;i<rows.length;i++){
            let seed = parseInt(key)+i
            let src = rows[i].image
            let animSpeed = rows[i].animSpeed
            let imgSize = rows[i].imgSize
            let speedDeviation = getRandomInt(-(animSpeed/2),0,seed)
            let xDeviation = getRandomInt(-100,+100,seed+2)
            let yDeviation = getRandomInt(-100,+100,seed+3)
            dogs.push(DogImageScrolling({deviation:{x:xDeviation,y:yDeviation}, key: i,alt: "", class: `${props.class} drop-shadow-xl`, src: src, style: {animationDuration:`${animSpeed+speedDeviation}ms`, height:`${imgSize}px`,width:`${imgSize}px`}}))
        }
        return (
            <div className={`absolute justify-around ${className}`} style={style} key={key}>
                {dogs}
            </div>
        );
        },[getRandomInt, props.class, props.direction])


    let getDogs = useCallback(()=>{
        let dogRows : JSX.Element[] = [];
        for(let dogImageRow in dogImageRows){
            dogRows.push(getDogRow(dogImageRows[dogImageRow],dogImageRow))
        }
        return dogRows
    },[dogImageRows, getDogRow])

    let content = useCallback(()=>{
        if(props.loading){
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
    },[getDogs, props.loading])

    let addDogRow = useCallback((images:string[],animSpeed:number,imgSize:number) => {
        let id:string = uid()
        let row: ImageRow[] = []
        for(let image in images){
            row.push({image:images[image],animSpeed:animSpeed,imgSize:imgSize})
        }
        updateDogImages({UID:id,newRow:row})
        setTimeout(()=>{
            updateDogImages({UID:id})
        },animSpeed)
    },[])



    const [dogImagesBuffer, setDogImageBuffer] = useState<string[]>([])

    let addNewDogRow = useCallback(()=>{
        if(props.filteredBreeds.length == 0){
            fetch(`api/dog/breeds/image/random/${props.dogCount}`)
                .then(res => res.json())
                .then(json => {
                    setDogImageBuffer(json['message'])
                    props.setLoading(false)})
        }
        else {
            for(let Breed in props.filteredBreeds){
                let eachBreed = Math.max(Math.floor(props.dogCount / props.filteredBreeds.length),1)
                fetch(`api/dog/breed/${props.filteredBreeds[Breed].value}/images/random/${eachBreed}`)
                    .then(res => res.json())
                    .then(json => {
                        setDogImageBuffer((prevState => prevState.concat(json["message"])))
                        props.setLoading(false)
                    })
            }
        }
    },[props])

    let [canLoop,setCanLoop] = useState(true)

    let loopAdd = useCallback(()=>{
        if(!canLoop){
            return
        }
        setCanLoop(false)
        setTimeout(()=>{
            addDogRow(dogImagesBuffer,props.animSpeed,props.imgSize)
            setCanLoop(true)
        }, Math.floor(1000))
        setDogImageBuffer([])
        addNewDogRow()
    },[canLoop, addNewDogRow, addDogRow, dogImagesBuffer, props.animSpeed, props.imgSize])




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