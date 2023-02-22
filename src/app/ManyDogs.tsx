'use client'
import DogImageScrolling from "@/app/DogImageScroll";
import React, {useEffect, useState} from "react";

interface Props {
    dogCount:number;
    class:string;
}

export default function ManyDogs(props:Props){

    const [dogImages, setDogImages] = useState([''])

    function getDogs(){
        let dogs : JSX.Element[] = [];
        for(let i=0;i<dogImages.length;i++){
            dogs.push(DogImageScrolling({alt: "", class: props.class, src: dogImages[i], style: {animationDuration:"30000ms",animationDelay:`${(i*4000).toString()}ms`}}))
        }
        return dogs;
    }

    useEffect(()=>{
        fetch(`api/breeds/image/random/${props.dogCount}`)
            .then(res => res.json())
            .then(json => {setDogImages(json['message'])})
    },[props.dogCount])

    return(
        <div className="fixed -left-72 w-full">
            <div className="absolute h-full -z-10">
                {getDogs()}
            </div>
        </div>
    )
}