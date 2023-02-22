import Image from "next/image";
import {CSSProperties, useEffect} from "react";

interface Props{
    src:string,
    alt:string
    class:string;
    style:CSSProperties
}
export default function DogImageScrolling(props:Props){
    return(
        <div className={props.class} style={props.style}>
            <Image src={props.src} alt={props.alt} width={200} height={200}/>
        </div>
    )
}