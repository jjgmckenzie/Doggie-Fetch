import Image from "next/image";
import {CSSProperties} from "react";

interface Vector{
    x:number,
    y:number
}

interface Props{
    key:number,
    src:string,
    alt:string
    class:string,
    style:CSSProperties
    deviation:Vector
}
export default function DogImageScrolling(props:Props){

    return(
        <div className={props.class} style={props.style} key={props.key}>
            <div style={{transform:`translateX(${props.deviation.x}px) translateY(${props.deviation.y}px)`}}>
                <Image src={props.src} alt={props.alt} width={200} height={200}/>
            </div>
        </div>
    )
}