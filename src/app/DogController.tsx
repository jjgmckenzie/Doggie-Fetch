import DirectionControl from "@/app/DirectionControl";
import {Dispatch, SetStateAction, useCallback, useState} from "react";
import DogDropDown, {Breed} from "@/app/DogDropDown";
import DogControlPanel from "@/app/DogControlPanel";
import Slider from 'rc-slider';
import 'rc-slider/assets/index.css';
import ClickAwayListener from 'react-click-away-listener';



interface Props{
    setDirection:Dispatch<SetStateAction<string>>
    setAnimSpeed:Dispatch<SetStateAction<number>>
    animSpeed:number
    imgSize:number
    setImgSize:Dispatch<SetStateAction<number>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    filteredBreeds:Breed[]
    direction:string
    loading:boolean
    breedList:Breed[]
}
export default function DogController(props:Props){
    const [optionsPoppedUp, setOptionsPoppedUp] = useState(false)
    const [uploadPoppedUp, setUploadPoppedUp] = useState(false)

    const FilterPanel = useCallback(()=>{
        if(optionsPoppedUp) {
            return (
                <div className="shadow-xl rounded-lg max-w-2xl w-[80vw] bg-white mb-2 p-4 mx-auto pointer-events-auto">
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
        if(uploadPoppedUp) {
            return (
                <div className="shadow-xl rounded-lg max-w-2xl w-[95vw] bg-white mb-2 px-4 pt-4 pb-2 mx-auto pointer-events-auto">
                    <h1 className="mx-auto text-xl sm:text-3xl text-center">Add your pooch to the Internet&apos;s biggest collection of <strong>open source dog pictures!</strong></h1>
                    <p className="my-2 px-1 leading-tight">The images available on this site are initially <i>fetched</i> from <a href="https://dog.ceo/api" className="text-blue-500 underline">Dog API</a>, who accept new members to their pack! You can upload your pictures here, and our bot will submit your good boy / girl to them on your behalf.</p>
                    <ul className="text-sm list list-disc leading-4 pl-7 pr-1 mb-2">
                        <li>Please ensure your photos are of a good quality and the dog is easily identifiable in the photo</li>
                        <li>Please ensure your photo features one dog only (although additional dogs can be in the background)</li>
                        <li>Photos must not include any human or any part of a human (GDPR)</li>
                    </ul>
                    <p className="px-1 text-xs sm:text-sm tracking-tight leading-3"><strong>Note: </strong>Fetch! uses AI vision to prevent misuse, which may flag your upload incorrectly. In this case, you can manually submit your photos as a <a className="text-blue-500 underline" href="https://github.com/jigsawpieces/dog-api-images#dog-api-images">GitHub pull request here.</a> </p>
                </div>
            )
        }
        return (
            <>
            </>
        )
    },[optionsPoppedUp, props, uploadPoppedUp])

    return(
        <div className="fixed bottom-0 right-0 left-0 w-min mx-auto z-10 mb-3 sm:mb-5">
            <ClickAwayListener onClickAway={()=>
            {setUploadPoppedUp(false)
            setOptionsPoppedUp(false)}}>
                <div className="pointer-events-none">
                    {FilterPanel()}
                    <div className="bg-white bg-opacity-25 rounded-lg backdrop-blur-sm shadow p-3 flex w-min mx-auto">
                        <DogControlPanel setOptionsPoppedUp={setOptionsPoppedUp} optionsPoppedUp={optionsPoppedUp} uploadPoppedUp={uploadPoppedUp} setUploadPoppedUp={setUploadPoppedUp}/>
                        <div className="mt-auto w-48">
                            <DirectionControl setDirection={props.setDirection}/>
                        </div>
                    </div>
                </div>
            </ClickAwayListener>
        </div>
    )
}