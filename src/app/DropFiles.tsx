import React, {useCallback, useEffect, useRef, useState} from 'react'

interface Props{
    isAcceptingFiles:boolean
    setFile: (file: File|null)=>void;
}

export default function  DropFiles(props:Props) {

    const [isDragActive,setIsDragActive] = useState(false)
    const blurRef = useRef<HTMLDivElement | null>(null);

    function dragEnter(e:Event){
        e.preventDefault()
    }

    function dragOver(e:Event){
        setIsDragActive(true)
        e.preventDefault()
    }

    function dragLeave(e:Event){
        e.preventDefault()
        setIsDragActive(false)
    }

    const dropFile = useCallback((e:DragEvent) => {
        e.preventDefault()
        console.log("dropped")
        console.log(e)
        console.log(props.isAcceptingFiles)
        if (props.isAcceptingFiles && e.dataTransfer && e.dataTransfer.items && e.dataTransfer.items.length > 0) {
            console.log("files found")
            let file = e.dataTransfer.items[0];
            if (file.kind === "file") {
                console.log("is a file")
                props.setFile(file.getAsFile())
            }
        }
        setIsDragActive(false)
    },[props])


    useEffect(()=>{
        document.addEventListener('drop',dropFile)
        return () => {
            document.removeEventListener('drop',dropFile)
        }

    },[dropFile])

    useEffect(()=>{
        document.addEventListener('dragenter',dragEnter)
        document.addEventListener('dragleave',dragLeave)
        document.addEventListener('dragover',dragOver)
        return () =>{
            document.removeEventListener('dragenter',dragEnter)
            document.removeEventListener('dragleave',dragLeave)
            document.removeEventListener('dragover',dragOver)
        }
    },[])

    useEffect(()=>{
        if(!blurRef.current){
            return
        }
        if(isDragActive && props.isAcceptingFiles){
            console.log("test blur")
           blurRef.current.style.setProperty('opacity','0.75')
        }
        else {
            console.log("test blur away")
            blurRef.current.style.setProperty('opacity','0')
        }
    },[isDragActive, props.isAcceptingFiles])

    return (
        <div className="absolute top-0 bottom-0 left-0 right-0 bg -z-10">
            <div ref={blurRef} className="bg-white w-full h-full transition-opacity duration-500"/>
        </div>
    )
}