'use client'
import {Dispatch, SetStateAction} from "react";

interface Props{
    setDirection:Dispatch<SetStateAction<string>>
}
export default function DirectionControl(props:Props){
    return(
        <div className="grid grid-cols-3 grid-rows-2 gap-y-3">
            <button className="col-span-3 ctrl-panel" onClick={()=> {props.setDirection("up")}}>↑</button>
            <button className="ctrl-panel" onClick={()=> {props.setDirection("left")}}>←</button>
            <button className="ctrl-panel" onClick={()=> {props.setDirection("down")}}>↓</button>
            <button className="ctrl-panel" onClick={()=> {props.setDirection("right")}}>→</button>
        </div>
    )
}