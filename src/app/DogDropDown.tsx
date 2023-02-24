'use client'
import Select from "react-dropdown-select";
import {Dispatch, SetStateAction} from "react";

export interface Breed{
    value:string,
    label:string,
}

interface Props{
    filteredBreeds:Breed[]
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    loading:boolean
    breedList:Breed[]
}


export default function DogDropDown(props:Props){

    return (
        <Select options={props.breedList} values={props.filteredBreeds} multi clearable dropdownPosition="top" placeholder="Select Breeds..." loading={props.loading} searchable onChange={(values)=>{props.setFilteredBreeds(values)}}/>
    )
}