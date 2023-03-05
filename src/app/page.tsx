'use client'
import ManyDogs from "@/app/ManyDogs";
import DogController from "@/app/DogController";
import {useCallback, useEffect, useState} from "react";
import {Breed} from "@/app/Breed";
import DropFiles from "@/app/DropFiles";

export default function Home() {
    const [filteredBreeds,setFilteredBreeds] = useState<Breed[]>([])
    const [direction,setDirection] = useState("right")
    const [animDirection,setAnimDirection] = useState("scrollRight")
    const [animSpeed,setAnimSpeed] = useState(30000)
    const [imgSize,setImgSize] = useState(300)
    const [breedList,setBreedList] = useState<Breed[]>([])
    const [loading,setLoading] = useState(false)
    const [image, setImage] = useState<string|null>(null);
    const [isAcceptingFiles, setIsAcceptingFiles] = useState(false)

    const setFile = useCallback((file:File|null)=>{
        if(file == null){
            setImage(null)
            return
        }
        let fileReader = new FileReader()
        fileReader.onload = (e) => {
            const result = e.target?.result
            if(result){
                setImage(result.toString())
            }
        }
        fileReader.readAsDataURL(file)
    },[])


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
        <main className="pt-4 select-none">
            <DropFiles isAcceptingFiles={isAcceptingFiles} setFile={setFile}/>
            <DogController setDirection={setDirection} setFilteredBreeds={setFilteredBreeds} filteredBreeds={filteredBreeds} direction={direction} setAnimSpeed={setAnimSpeed} animSpeed={animSpeed} imgSize={imgSize} setImgSize={setImgSize} loading={loading} breedList={breedList} image={image} setFile={setFile} setIsAcceptingFiles={setIsAcceptingFiles}/>
            <ManyDogs dogCount={6} class={animDirection} filteredBreeds={filteredBreeds} direction={direction} animSpeed={animSpeed} imgSize={imgSize} loading={loading} setLoading={setLoading}/>
            <span className="mx-auto left-0 right-0 bottom-0 absolute w-min" id="confettiReward"/>
        </main>
    )
}
