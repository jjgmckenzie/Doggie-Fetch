'use client'
import Select from "react-dropdown-select";
import {Dispatch, SetStateAction, useEffect, useState} from "react";

export interface Breed{
    value:string,
    label:string,
}

interface Props{
    filteredBreeds:Breed[]
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
}


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

export default function DogDropDown(props:Props){
    const [breedList,setBreedList] = useState<Breed[]>([])
    const [loading,setLoading] = useState(true)
    useEffect((()=>{
        fetch('/api/dog/breeds/list/all')
            .then(res => res.json())
            .then(res => res["message"])
            .then(res => ParseBreedJson(res))
            .then(res=> {
                setBreedList(res)
                setLoading(false)})
            })
    ,[])
    return (
        <Select options={breedList} values={props.filteredBreeds} multi clearable dropdownPosition="top" loading={loading} searchable onChange={(values)=>{props.setFilteredBreeds(values)}}/>
    )
}