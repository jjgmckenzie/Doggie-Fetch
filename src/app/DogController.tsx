import DirectionControl from "@/app/DirectionControl";
import {Dispatch, SetStateAction, useCallback, useState} from "react";
import DogDropDown, {Breed} from "@/app/DogDropDown";
import DogControlPanel from "@/app/DogControlPanel";
import Slider from 'rc-slider';
import 'rc-slider/assets/index.css';


interface Props{
    setDirection:Dispatch<SetStateAction<string>>
    setAnimSpeed:Dispatch<SetStateAction<number>>
    animSpeed:number
    imgSize:number
    setImgSize:Dispatch<SetStateAction<number>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    filteredBreeds:Breed[]
    direction:string
}
export default function DogController(props:Props){
    const [filterPoppedUp, setFilterPoppedUp] = useState(false)

    const FilterPanel = useCallback(()=>{
        if(filterPoppedUp) {
            return (
                <div className="shadow-xl rounded-lg max-w-2xl w-[80vw] bg-white mb-2 p-4 mx-auto">
                    <DogDropDown {...props}/>
                    <div className="flex mt-2">
                        <label className=" text-sm mr-4 text-center">Speed:</label>
                        <Slider min={(1/50_000)} max={(1/10_000)} defaultValue={(1/props.animSpeed)} step={(1/100_000)} className={"my-auto"} onChange={value => {props.setAnimSpeed(1/(value as number))}}/>
                        <label className=" text-sm mx-4 text-center">Size:</label>
                        <Slider min={100} max={500} defaultValue={props.imgSize} className={"my-auto"} onChange={value => {props.setImgSize(value as number)}}/>
                    </div>
                </div>
            )
        }
        return (
            <>
            </>
        )
    },[filterPoppedUp, props])

    return(
        <div className="fixed bottom-0 right-0 left-0 w-min mx-auto z-10 mb-3 sm:mb-5">
            {FilterPanel()}
            <div className="bg-white bg-opacity-25 rounded-lg backdrop-blur-sm shadow p-3 flex w-min mx-auto">
                <DogControlPanel setFilterPoppedUp={setFilterPoppedUp}/>
                <div className="mt-auto w-60">
                    <DirectionControl setDirection={props.setDirection}/>
                </div>
            </div>
        </div>
    )
}