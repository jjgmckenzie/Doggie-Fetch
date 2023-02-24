'use client'
import ManyDogs from "@/app/ManyDogs";
import DogController from "@/app/DogController";
import {useEffect, useState} from "react";
import {Breed} from "@/app/Breed";

export default function Home() {
    const [filteredBreeds,setFilteredBreeds] = useState<Breed[]>([])
    const [direction,setDirection] = useState("right")
    const [animDirection,setAnimDirection] = useState("scrollRight")
    const [animSpeed,setAnimSpeed] = useState(30000)
    const [imgSize,setImgSize] = useState(300)
    const [breedList,setBreedList] = useState<Breed[]>([])
    const [loading,setLoading] = useState(true)
    function ParseBreedJson(breedJson:[string:[string]]) : Breed[] {

        let parsedBreeds: Breed[] = []
        function pushBreed(breed:string, subBreed?:string){
            function capitalizeFirstLetter(string:string) {
                return string.charAt(0).toUpperCase() + string.slice(1);
            }

            let value = `${breed}`
            let label = capitalizeFirstLetter(breed)
            if(subBreed){
                label = `${capitalizeFirstLetter(subBreed)} ${label}`
                value = `${value}/${subBreed}`
            }
            parsedBreeds.push({value:value,label:label})
        }
        for(let breed in breedJson){
            if(breedJson[breed].length > 0){
                for(let subBreed in breedJson[breed]){
                    pushBreed(breed, breedJson[breed][subBreed])
                }
            }
            else{
                pushBreed(breed)
            }
        }
        return parsedBreeds
    }
    useEffect((()=>{
            fetch('/api/dog/breeds/list/all')
                .then(res => res.json())
                .then(res => res["message"])
                .then(res => ParseBreedJson(res))
                .then(res=> {
                    setBreedList(res)})
        })
        ,[])
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
            <DogController setDirection={setDirection} setFilteredBreeds={setFilteredBreeds} filteredBreeds={filteredBreeds} direction={direction} setAnimSpeed={setAnimSpeed} animSpeed={animSpeed} imgSize={imgSize} setImgSize={setImgSize} loading={loading} breedList={breedList}/>
            <ManyDogs dogCount={6} class={animDirection} filteredBreeds={filteredBreeds} direction={direction} animSpeed={animSpeed} imgSize={imgSize} loading={loading} setLoading={setLoading}/>
        </main>
    )
}
