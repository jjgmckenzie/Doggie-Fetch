'use client'
import {Dispatch, SetStateAction} from "react";

interface Props{
    setDirection:Dispatch<SetStateAction<string>>
}
export default function DirectionControl(props:Props){
    return(
        <div className="grid grid-cols-3 grid-rows-2 gap-3">
            <button className="col-span-3 dirButton" onClick={()=> {props.setDirection("up")}}>↑</button>
            <button className="dirButton" onClick={()=> {props.setDirection("left")}}>←</button>
            <button className="dirButton" onClick={()=> {props.setDirection("down")}}>↓</button>
            <button className="dirButton" onClick={()=> {props.setDirection("right")}}>→</button>
        </div>
    )
}