import DirectionControl from "@/app/DirectionControl";
import {Dispatch, SetStateAction} from "react";
import DogDropDown, {Breed} from "@/app/DogDropDown";

interface Props{
    setDirection:Dispatch<SetStateAction<string>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    filteredBreeds:Breed[]
    direction:string
}
export default function DogController(props:Props){
    return(
        <div className="bg-white rounded-lg shadow p-3 fixed bottom-0 right-0 left-0 max-w-4xl mx-auto z-10 mb-5 w-[95vw] flex">
            <div className="mr-auto my-auto max-w-[60%] max-h-[32]">
                <DogDropDown setFilteredBreeds={props.setFilteredBreeds} filteredBreeds={props.filteredBreeds}/>
            </div>
            <div className="ml-2 mt-auto">
                <DirectionControl setDirection={props.setDirection}/>
            </div>
        </div>
    )
}