'use client'
import ManyDogs from "@/app/ManyDogs";
import DogController from "@/app/DogController";
import {useEffect, useState} from "react";

export default function Home() {
    const [direction,setDirection] = useState("right")
    const [animDirection,setAnimDirection] = useState("scrollRight")
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
            <DogController setDirection={setDirection}/>
            <ManyDogs dogCount={50} class={animDirection}/>
        </main>
    )
}
