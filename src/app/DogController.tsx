import DirectionControl from "@/app/DirectionControl";
import {Dispatch, SetStateAction} from "react";

interface Props{
    setDirection:Dispatch<SetStateAction<string>>
}
export default function DogController(props:Props){
    return(
        <div className="bg-white rounded-lg shadow p-3 fixed bottom-0 right-0 left-0 max-w-4xl mx-auto z-10 mb-5 w-[80vw]">
            <DirectionControl setDirection={props.setDirection}/>
        </div>
    )
}