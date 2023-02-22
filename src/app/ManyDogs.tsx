'use client'
import DogImageScrolling from "@/app/DogImageScroll";
import React, {useCallback, useEffect, useState} from "react";
import {Breed} from "@/app/DogDropDown";

interface Props {
    dogCount:number,
    class:string,
    filteredBreeds:Breed[],
}

export default function ManyDogs(props:Props){

    const [dogImages, setDogImages] = useState([''])

    function getRandomInt(min:number, max:number) {
        min = Math.ceil(min);
        max = Math.floor(max);
        return Math.floor(Math.random() * (max - min + 1) + min); // The maximum is inclusive and the minimum is inclusive
    }

    let getDogs = useCallback(()=>{
        let dogs : JSX.Element[] = [];
        for(let i=0;i<dogImages.length;i++){
            let speedDeviation = getRandomInt(-10000,10000)
            let delayDeviation = getRandomInt(-1000,10000)
            let xDeviation = getRandomInt(-100,+100)
            let yDeviation = getRandomInt(-100,+100)
            dogs.push(DogImageScrolling({deviation:{x:xDeviation,y:yDeviation}, key: i,alt: "", class: `${props.class} drop-shadow-lg`, src: dogImages[i], style: {animationDuration:`${30000+speedDeviation}ms`,animationDelay:`${(-i*4000+delayDeviation).toString()}ms`}}))
        }
        return dogs;
        },[dogImages, props.class])


    useEffect(()=>{

    },[])

    useEffect(()=>{
        console.log("updating dog images")
        if(props.filteredBreeds.length == 0){
            fetch(`api/dog/breeds/image/random/${props.dogCount}`)
                .then(res => res.json())
                .then(json => {setDogImages(json['message'])})
        }
        else {
            let eachBreed = Math.floor(props.dogCount / props.filteredBreeds.length)
            setDogImages([])
            for(let Breed in props.filteredBreeds){
                console.log(`api/dog/breed/${props.filteredBreeds[Breed].value}/images/random/${eachBreed}`)
                fetch(`api/dog/breed/${props.filteredBreeds[Breed].value}/images/random/${eachBreed}`)
                    .then(res => res.json())
                    .then(json => {
                        console.log(json["message"])
                        setDogImages(prevState => (prevState.concat(json["message"])))
                    })
            }
        }
    },[props.dogCount, props.filteredBreeds])


    return(
        <div className="fixed w-full h-full -z-10">
            <div className="absolute grid w-full h-full top-0 left-0 grid-cols-8 grid-rows-8 gap-2">
                {getDogs()}
            </div>
        </div>
    )
}