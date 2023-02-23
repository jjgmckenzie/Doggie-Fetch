import DirectionControl from "@/app/DirectionControl";
import {Dispatch, SetStateAction, useCallback, useState} from "react";
import DogDropDown, {Breed} from "@/app/DogDropDown";
import DogControlPanel from "@/app/DogControlPanel";

interface Props{
    setDirection:Dispatch<SetStateAction<string>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    filteredBreeds:Breed[]
    direction:string
}
export default function DogController(props:Props){
    const [filterPoppedUp, setFilterPoppedUp] = useState(false)

    const FilterPanel = useCallback(()=>{
        if(filterPoppedUp) {
            return (
                <div className="shadow rounded-lg max-w-3xl w-[90vw] bg-white mb-4 p-1 mx-auto">
                    <DogDropDown {...props}/>
                </div>
            )
        }
        return (
            <>
            </>
        )
    },[filterPoppedUp, props])

    return(
        <div className="fixed bottom-0 right-0 left-0 max-w-4xl mx-auto z-10 mb-5 w-[95vw]">
            {FilterPanel()}
            <div className="bg-white rounded-lg shadow p-3 flex">
                <div className="mr-auto my-auto max-w-[60%] max-h-[32]">
                    <DogControlPanel setFilterPoppedUp={setFilterPoppedUp}/>
                </div>
                <div className="ml-2 mt-auto">
                    <DirectionControl setDirection={props.setDirection}/>
                </div>
            </div>
        </div>
    )
}