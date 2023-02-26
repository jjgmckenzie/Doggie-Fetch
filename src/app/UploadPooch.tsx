"use client"
import Select from "react-dropdown-select";
import {isMobile} from "react-device-detect";
import {Breed} from "@/app/Breed";
import {Dispatch, SetStateAction, useCallback, useEffect, useRef, useState} from "react";

interface Props {
    breedList:Breed[]
    breedUploaded:Breed[]
    setBreedUploaded:Dispatch<SetStateAction<Breed[]>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    loading:boolean
    image:string|null
    setFile: (file: File|null)=>void;
    setIsAcceptingFiles:Dispatch<SetStateAction<boolean>>
}

export default function UploadPooch(props:Props){
    const inputRef = useRef<HTMLInputElement | null>(null);
    const [checked,setChecked] = useState(false)
    const [acceptedTerms,setAcceptedTerms] = useState(false)
    const TermsAndConditions = useCallback(() => (
            <>
                <p className="mb-1 px-1 leading-tight text-sm sm:text-base">The images available on this site are initially <i>fetched</i> from <a href="https://dog.ceo/api" target="_blank" className="text-blue-500 underline" rel="noreferrer">Dog API</a>, who accept new members to their pack! You can upload your photos here, and our bot will resize and submit your good boy / girl for approval on your behalf.</p>
                <strong>Terms & Conditions:</strong>
                <ul className="text-sm sm:text-base list list-disc leading-4 pl-7 pr-1 mb-2">
                    <li>Photos must be of a good quality with an easily identifiable dog</li>
                    <li>You must have the rights to submit and release the photo to the public domain</li>
                    <li>Photos must feature one prominent dog only <i>(although additional dogs may be in the background)</i></li>
                    <li>Photos must not include any human or any part of a human (GDPR)</li>
                    <li>Your photos will be made available through the <a className="text-blue-500 underline" href="https://dog.ceo/dog-api/documentation/" target="_blank" rel="noreferrer">API endpoints</a></li>
                </ul>
                <div className="text-sm sm:text-base flex">
                    <label className="pl-2 my-auto">
                        <input type="checkbox" checked={checked} onChange={()=>{setChecked(prevState=> !prevState)}}/>
                        <span className="pl-2">Accept Terms and Conditions</span>
                    </label>
                    <button disabled={!checked} className="ml-auto text-base sm:text-lg bg-gray-100 enabled:bg-blue-600 enabled:text-white text-gray-400  font-bold px-1 py-0.5 rounded-md " onClick={()=>{setAcceptedTerms(true)}}>Next →</button>
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
                    if(e.target.files){
                        props.setFile(e.target.files[0])
                        e.target.value = ""
                    }
                }}
            />
            <button className="bg-gray-200 border- border-gray-500 rounded p-1 z-10" onClick={() => {inputRef.current?.click()}}>Select Photo</button>
        </div> )
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
                    <input className="bg-transparent border-2 p-1 w-full"/>
                </div>
                <div className="mx-4">
                    <label><strong className="text-sm sm:text-base">Breed:</strong> <i className="text-xs sm:text-sm">(We love mixed breeds too! Select &apos;Mix&apos;)</i></label>
                    <Select options={props.breedList} values={props.breedUploaded} searchable={!isMobile} dropdownHandle={false} dropdownPosition="top" placeholder="Select Breed (If a mix, select 'Mix')" loading={props.loading}  onChange={(values)=>{
                        props.setBreedUploaded(values)
                        props.setFilteredBreeds(values)}} />
                </div>
                <div className="flex justify-between mt-1 sm:mt-2">
                    <button className="sm:text-lg text-blue-600 border-2 font-bold px-1 rounded-md " onClick={()=>{setAcceptedTerms(false)}}>Back</button>
                    <button className="sm:text-lg bg-blue-600 text-white font-bold px-1 rounded-md " onClick={()=>{}}>Submit</button>
                </div>
            </div>
        )
    },[FileBox, props, uploadButton])

    const getBody = useCallback(()=>{
       if(!acceptedTerms){
           return TermsAndConditions()
       }
       return submitDogForm()
        }
       ,[TermsAndConditions, acceptedTerms, submitDogForm])



    let { setIsAcceptingFiles } = props
    useEffect(()=>{
        setIsAcceptingFiles(checked)
    },[checked, setIsAcceptingFiles])

    return (
        <div className="shadow-xl rounded-lg max-w-2xl w-[95vw] glass-bg mb-2 p-2 sm:px-4 mx-auto pointer-events-auto">
            <h2 className="mx-auto text-xl sm:text-3xl text-center leading-tight mb-1 sm:mb-2">Add your pooch to the Internet&apos;s biggest collection of <strong>open source dog pictures!</strong></h2>
            {getBody()}
            <p className="mt-2 px-1 text-xs sm:text-sm tracking-tight leading-3 text-center"><strong>Note: </strong>Go Fetch! uses AI vision to prevent misuse, which may incorrectly prevent your upload. In this case, or if your dog&apos;s breed is not listed, you can manually submit your photos as a <a className="text-blue-500 underline" href="https://github.com/jigsawpieces/dog-api-images#dog-api-images" target="_blank" rel="noreferrer" >GitHub pull request here.</a> </p>
        </div>
    )
}