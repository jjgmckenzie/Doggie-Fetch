'use client'
import ManyDogs from "@/app/ManyDogs";
import DogController from "@/app/DogController";
import {useEffect, useState} from "react";
import {Breed} from "@/app/DogDropDown";

export default function Home() {
    const [filteredBreeds,setFilteredBreeds] = useState<Breed[]>([])
    const [direction,setDirection] = useState("right")
    const [animDirection,setAnimDirection] = useState("scrollRight")
    const [animSpeed,setAnimSpeed] = useState(30000)
    const [imgSize,setImgSize] = useState(300)
    useEffect(()=>{
        switch (direction) {
            case "up":{
                setAnimDirection("scrollUp")
                return;
            }
            case "down":{
                setAnimDirection("scrollDown")
                return;
            }
            case "left":{
                setAnimDirection("scrollLeft")
                return;
            }
            case "right":{
                setAnimDirection("scrollRight")
                return;
            }
        }
    },[direction])

    return (
        <main className="pt-4">
            <DogController setDirection={setDirection} setFilteredBreeds={setFilteredBreeds} filteredBreeds={filteredBreeds} direction={direction} setAnimSpeed={setAnimSpeed} animSpeed={animSpeed} imgSize={imgSize} setImgSize={setImgSize}/>
            <ManyDogs dogCount={6} class={animDirection} filteredBreeds={filteredBreeds} direction={direction} animSpeed={animSpeed} imgSize={imgSize}/>
        </main>
    )
}
