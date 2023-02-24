`use client`
import {Dispatch, SetStateAction} from "react";


interface Props{
    setFilterPoppedUp:Dispatch<SetStateAction<boolean>>
}

export default function DogControlPanel (props:Props) {
    return (
            <button className="ctrl-panel text-2xl p-2 min-w-fit my-auto" onClick={()=>{props.setFilterPoppedUp(prevState => !prevState)}}>Options</button>
    )
}