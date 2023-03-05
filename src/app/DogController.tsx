import DirectionControl from "@/app/DirectionControl";
import {Dispatch, SetStateAction, useCallback, useState} from "react";
import {Breed} from "@/app/Breed";
import DogControlPanel from "@/app/DogControlPanel";
import Slider from 'rc-slider';
import 'rc-slider/assets/index.css';
import ClickAwayListener from 'react-click-away-listener';
import Select from "react-dropdown-select";
import {isMobile} from 'react-device-detect';
import UploadPooch from "@/app/UploadPooch";



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
    setFile: (file: File|null)=>void;
    image:string|null
    setIsAcceptingFiles: Dispatch<SetStateAction<boolean>>
}
export default function DogController(props:Props){
    const [optionsPoppedUp, setOptionsPoppedUp] = useState(false)
    const [uploadingPooch,setUploadingPooch] = useState(false)
    const [dogName,setDogName] = useState("")
    const [uploadPoppedUp, setUploadPoppedUp] = useState(false)
    const [breedUploaded,setBreedUploaded] = useState<Breed[]>([])
    const [response,setResponse] = useState(0)

    const FilterPanel = useCallback(()=>{
        if(optionsPoppedUp) {
            return (
                <div className="shadow-xl rounded-lg max-w-2xl w-[80vw] glass-bg mb-2 p-4 mx-auto pointer-events-auto">
                    <Select options={props.breedList} values={props.filteredBreeds} multi clearable dropdownPosition="top" placeholder="Select Breeds..." loading={props.loading} searchable={!isMobile} onChange={(values)=>{props.setFilteredBreeds(values)}}/>
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
                <UploadPooch breedList={props.breedList} breedUploaded={breedUploaded} loading={props.loading} setBreedUploaded={setBreedUploaded} setFilteredBreeds={props.setFilteredBreeds} image={props.image} setFile={props.setFile} setIsAcceptingFiles={props.setIsAcceptingFiles} loadingPooch={uploadingPooch} setIsLoadingPooch={setUploadingPooch} dogName={dogName} setDogName={setDogName} response={response} setResponse={setResponse}/>
            )
        }
        return (
            <>
            </>
        )
    },[breedUploaded, dogName, optionsPoppedUp, props, response, uploadPoppedUp, uploadingPooch])
    const ClickAway = useCallback(()=>{
        if(!uploadingPooch){
            setUploadPoppedUp(false)
            setOptionsPoppedUp(false)
        }
    },[uploadingPooch])

    return(
        <div className="fixed bottom-0 right-0 left-0 w-min mx-auto z-10 mb-1 sm:mb-5">
            <ClickAwayListener onClickAway={ClickAway}>
                <div className="pointer-events-none">
                    {FilterPanel()}
                    <div className="bg-white bg-opacity-25 rounded-lg backdrop-blur-sm shadow p-2 sm:p-4 flex w-min mx-auto">
                        <DogControlPanel setOptionsPoppedUp={setOptionsPoppedUp} optionsPoppedUp={optionsPoppedUp} uploadPoppedUp={uploadPoppedUp} setUploadPoppedUp={setUploadPoppedUp} />
                        <div className="mt-auto w-48">
                            <DirectionControl setDirection={props.setDirection}/>
                        </div>
                    </div>
                </div>
            </ClickAwayListener>
        </div>
    )
}