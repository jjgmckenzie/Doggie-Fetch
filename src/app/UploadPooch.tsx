"use client"
import Select from "react-dropdown-select";
import {isMobile} from "react-device-detect";
import {Breed} from "@/app/Breed";
import React, {Dispatch, SetStateAction, useCallback, useEffect, useRef, useState} from "react";
import SyncLoader from "react-spinners/SyncLoader";
import {useReward} from "react-rewards";

interface Props {
    breedList:Breed[]
    breedUploaded:Breed[]
    setBreedUploaded:Dispatch<SetStateAction<Breed[]>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    loading:boolean
    image:string|null
    setFile: (file: File|null)=>void;
    setIsAcceptingFiles:Dispatch<SetStateAction<boolean>>
    loadingPooch:boolean
    setIsLoadingPooch:Dispatch<SetStateAction<boolean>>
    dogName: string
    setDogName: Dispatch<SetStateAction<string>>
    response: number
    setResponse: Dispatch<SetStateAction<number>>
}

export default function UploadPooch(props:Props){
    const inputRef = useRef<HTMLInputElement | null>(null);
    const [checked,setChecked] = useState(false)
    const [link,setLink] = useState("")
    const { reward } = useReward('confettiReward', 'confetti', {spread:isMobile? 50 : 120,elementCount:isMobile ? 50 : 200,elementSize:16,startVelocity:isMobile? 20 : 30,decay:0.98,lifetime:600,zIndex:20});
    const [acceptedTerms,setAcceptedTerms] = useState(false)
    const TermsAndConditions = useCallback(() => (
            <>
                <p className="mb-1 px-1 leading-tight text-sm sm:text-base">The images available on this site are initially <i>fetched</i> from <a href="https://dog.ceo/api" target="_blank" className="text-blue-500 underline" rel="noreferrer">Dog API</a>, who accept new members to their pack! You can upload your photos here, and our bot will resize and submit your good boy / girl for approval on your behalf.</p>
                <strong>Terms & Conditions:</strong>
                <ul className="text-sm sm:text-base list list-disc leading-4 pl-7 pr-1">
                    <li>Photos must be of a good quality with an easily identifiable dog</li>
                    <li>You must have the rights to submit and release the photo to the public domain</li>
                    <li>Photos must feature one prominent dog only <i>(although additional dogs may be in the background)</i></li>
                    <li>Photos must not include any human or any part of a human (GDPR)</li>
                    <li>Your photos will be made available through the <a className="text-blue-500 underline" href="https://dog.ceo/dog-api/documentation/" target="_blank" rel="noreferrer">API endpoints</a></li>
                </ul>
                <div className="text-sm sm:text-base flex">
                    <label className="pl-3 justify-content-center flex pb-2 pt-1">
                            <input type="checkbox" checked={checked} onChange={()=>{setChecked(prevState=> !prevState)}}/>
                            <span className="pl-2 my-auto">Accept Terms and Conditions</span>
                    </label>
                    <button disabled={!checked} className="ml-auto text-base sm:text-lg bg-gray-100 enabled:bg-blue-600 enabled:text-white text-gray-400  font-bold px-1 py-0.5 rounded-md mb-2 mt-1" onClick={()=>{setAcceptedTerms(true)}}>Next →</button>
                </div>
            </>
        ),[checked])


    const dragAndDrop = useCallback(()=>{
        if(!isMobile){
            return <span className="m-auto w-fit font-bold text-gray-500 text-2xl">Drag and drop photo here</span>
        }
    },[])


    const displayImage = useCallback((image:string)=> {
        return (<div className="mx-auto relative">
            <button className="rounded-full absolute bg-red-600 text-white px-1 text-xs right-0 translate-x-1.5 -translate-y-1.5 border-2 border-black font-extrabold pointer-events-auto" onClick={()=>{props.setFile(null)}}>X</button>
            {/* eslint-disable-next-line @next/next/no-img-element  -- we do not care to optimize a clientside image*/}
            <img className="min-h-[5rem] max-h-full max-w-full object-scale-down shadow-lg mx-auto" src={image} alt=""/>
        </div>)
    },[props])

    const displayBox = useCallback(()=>{
        let body;
        if(props.image){
            body = displayImage(props.image)
        }
        else{
            body = dragAndDrop()
        }
        return(
            <div className="absolute top-0 bottom-0 left-0 right-0 w-full h-full flex pointer-events-none">
                {body}
            </div>
        )
    },[displayImage, dragAndDrop, props.image])



    const FileBox = useCallback(()=>{
        return (
            <div className="flex h-full w-full relative">
                {displayBox()}
            </div>
        )
    },[displayBox])

    const uploadButton = useCallback(()=>{
        return ( <div className="text-center my-1">
            <input
                ref={inputRef}
                type="file"
                accept="image/*"
                style={{ display: 'none' }}
                onChange={(e) => {
                    if(e.target.files != null && e.target.files.length > 0){
                        props.setFile(e.target.files[0])
                        e.target.value = ""
                    }
                }}
            />
            <button className="bg-gray-200 border- border-gray-500 rounded p-1 z-10" onClick={() => {inputRef.current?.click()}}>Select Photo</button>
        </div> )
    },[props])

    const postDog = useCallback(()=>{
            fetch("/upload",{
                method: "POST",
                headers: {
                    "Content-Type" : "application/json",
                },
                body:JSON.stringify({
                    name:props.dogName,
                    breed:props.breedUploaded[0].value.replace("/","-"),
                    image:props.image,
                })
            }).then(res => {
                if(res.status == 202){
                    reward()
                    setTimeout(()=>{
                        reward()
                    },5000)
                    setTimeout(()=>{
                        reward()
                    },10000)

                    if(res.body){
                        res.text()
                            .then(pullLink => {
                                setLink(pullLink.slice(1,-1))
                            })
                    }
                }

                props.setResponse(res.status)
                setTimeout(()=>{
                    props.setIsLoadingPooch(false)
                },1000)
            })
            props.setIsLoadingPooch(true)
        }
        ,[props, reward])

    const clearForm = useCallback(()=>{
        props.setFile(null)
        let breedList = props.breedList
        props.setBreedUploaded([])
        props.setFilteredBreeds(breedList)
        props.setDogName("")
        props.setResponse(0)
    },[props])


    const clearImage = useCallback(()=>{
        props.setFile(null)
        props.setResponse(0)
    },[props])

    const submitDogForm = useCallback( () => {
        return (
            <div className="flex flex-col">
                <div className="flex-grow max-h-[15vh] h-72">
                    {FileBox()}
                </div>
                {uploadButton()}
                <div className="mx-4">
                    <label><strong className="text-sm sm:text-base">Name:</strong> <i className="text-xs sm:text-sm">(The dog&apos;s name, not yours!)</i></label>
                    <input className="bg-transparent border-2 p-1 w-full" defaultValue={props.dogName} onChange={(e)=>{props.setDogName(e.target.value)}}/>
                </div>
                <div className="mx-4">
                    <label><strong className="text-sm sm:text-base">Breed:</strong> <i className="text-xs sm:text-sm">(We love mixed breeds too! Select &apos;Mix&apos;)</i></label>
                    <Select options={props.breedList} values={props.breedUploaded} searchable={!isMobile} dropdownHandle={false} dropdownPosition="top" placeholder="Select Breed (If a mix, select 'Mix')" loading={props.loading}  onChange={(values)=>{
                        props.setBreedUploaded(values)
                        props.setFilteredBreeds(values)}} />
                </div>
                <div className="flex justify-between mt-1 sm:mt-2">
                    <button className="sm:text-lg text-blue-600 border-2 font-bold px-1 rounded-md " onClick={()=>{setAcceptedTerms(false)}}>Back</button>
                    <button className="sm:text-lg bg-blue-600 text-white font-bold px-1 rounded-md " onClick={postDog}>Submit</button>
                </div>
            </div>
        )
    },[FileBox, postDog, props, uploadButton])

    const uploadingPooch = useCallback( ()=>{
        let loadingLine = `Welcoming ${props.dogName} to the pack! 🐶`
        return (
            <div className="h-[30vh] mx-auto w-fit flex flex-col mb-2">
                { props.response == 0 && (
                   <>
                       <SyncLoader className="m-auto w-fit"/>
                       <span className="text-xl sm:text-2xl">{loadingLine}</span>
                   </>
                    )
                }
                { props.response == 202 && (
                    <>
                        <span className="text-4xl sm:text-5xl font-bold mx-auto mt-auto mb-8"> 🎉 Success! 🎉</span>
                        <span className="mx-auto text-center sm:text-lg sm:px-24">Thank you for sharing {props.dogName} with us! We can&apos;t wait to welcome our newest pack member</span>
                        <a className="mx-auto text-center text-sm sm:text-base text-blue-600 underline font-bold mb-auto" href={link} rel="noreferrer" target="_blank">Feel free to tell us more about them here!</a>
                        <div className="flex"><button disabled={props.loadingPooch} onClick={()=>clearForm()} className="ml-auto rounded bg-blue-600 disabled:bg-gray-100 disabled:text-gray-400 text-white border-1 font-bold p-1 sm:text-lg mr-1 mb-1 transition-colors duration-1000">Upload Another</button></div>
                    </>
                )}
                { props.response == 412 && (
                    <>
                        <span className="text-4xl sm:text-5xl font-bold mx-auto mt-auto mb-4">🤖 Try again 🤖</span>
                        <span className="mx-auto text-center sm:text-lg sm:px-24 mb-4">We&apos;re sorry, but our robot doesn&apos;t like this photo. This could be due to an error, due to low quality, or due to a human in it!  </span>
                        <div className="flex"><button disabled={props.loadingPooch} onClick={()=>clearImage()} className="ml-auto rounded bg-blue-600 disabled:bg-gray-100 disabled:text-gray-400 text-white border-1 font-bold p-1 sm:text-lg mr-1 mb-2 transition-colors duration-1000">Try Another Picture</button></div>
                    </>
                )
                }

                { props.response != 0 && props.response != 412 && props.response != 202 && (
                    <>
                        <span className="text-4xl sm:text-5xl font-bold mx-auto mt-auto mb-4">🦮 Uh Oh! 🦮</span>
                        <span className="mx-auto text-center sm:text-lg sm:px-24 mb-4"> Something went wrong. The server may be down, under high load, or the image file may be an unreadable format. If this persists; email me or make a GitHub issue.  </span>
                        <div className="flex"><button disabled={props.loadingPooch} onClick={()=>props.setResponse(0)} className="ml-auto rounded bg-blue-600 disabled:bg-gray-100 disabled:text-gray-400 text-white border-1 font-bold p-1 sm:text-lg mr-1 mb-2 transition-colors duration-1000">Try Again</button></div>
                    </>
                )
                }
            </div>
        )
    },[clearForm, clearImage, link, props])

    const getBody = useCallback(()=>{
        if(props.loadingPooch || props.response != 0){
            return uploadingPooch()
        }
       if(!acceptedTerms){
           return TermsAndConditions()
       }
       return submitDogForm()
        }
       ,[TermsAndConditions, acceptedTerms, props.loadingPooch, props.response, submitDogForm, uploadingPooch])



    let { setIsAcceptingFiles } = props
    useEffect(()=>{
        setIsAcceptingFiles(checked)
    },[checked, setIsAcceptingFiles])

    return (
        <div className="shadow-xl rounded-lg max-w-2xl w-[95vw] glass-bg mb-2 sm:mb-4 p-2 sm:px-4 mx-auto pointer-events-auto">
            <h2 className="mx-auto text-xl sm:text-3xl text-center leading-tight mb-1 sm:mb-2">Add your pooch to the Internet&apos;s biggest collection of <strong>open source dog pictures!</strong></h2>
            {getBody()}
            <p className="px-1 text-xs sm:text-sm tracking-tight leading-3 text-center"><strong>Note: </strong>Go Fetch! uses AI vision to prevent misuse, which may incorrectly prevent your upload. In this case, or if your dog&apos;s breed is not listed, you can manually submit your photos as a <a className="text-blue-500 underline" href="https://github.com/jigsawpieces/dog-api-images#dog-api-images" target="_blank" rel="noreferrer" >GitHub pull request here.</a> </p>
        </div>
    )
}