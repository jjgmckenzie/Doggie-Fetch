import Select from "react-dropdown-select";
import {isMobile} from "react-device-detect";
import {Breed} from "@/app/Breed";
import {Dispatch, SetStateAction, useCallback, useState} from "react";

interface Props {
    breedList:Breed[]
    breedUploaded:Breed[]
    setBreedUploaded:Dispatch<SetStateAction<Breed[]>>
    setFilteredBreeds:Dispatch<SetStateAction<Breed[]>>
    loading:boolean
}

export default function UploadPooch(props:Props){

    const [checked,setChecked] = useState(false)
    const [acceptedTerms,setAcceptedTerms] = useState(false)
    const TermsAndConditions = useCallback(() => (
            <>
                <p className="mb-1 px-1 leading-tight text-sm sm:text-base">The images available on this site are initially <i>fetched</i> from <a href="https://dog.ceo/api" target="_blank" className="text-blue-500 underline" rel="noreferrer">Dog API</a>, who accept new members to their pack! You can upload your photos here, and our bot will resize and submit your good boy / girl for approval on your behalf.</p>
                <strong>Terms & Conditions:</strong>
                <ul className="text-sm list list-disc leading-4 pl-7 pr-1 mb-2">
                    <li>Photos must be of a good quality with an easily identifiable dog</li>
                    <li>You must have the rights to submit and release the photo to the public domain</li>
                    <li>Photos must feature one prominent dog only <i>(although additional dogs may be in the background)</i></li>
                    <li>Photos must not include any human or any part of a human (GDPR)</li>
                    <li>Your photos will be made available through the <a className="text-blue-500 underline" href="https://dog.ceo/dog-api/documentation/" target="_blank" rel="noreferrer">API endpoints</a></li>
                </ul>
                <div className="text-sm flex">
                    <label className="pl-2 my-auto">
                        <input type="checkbox" checked={checked} onChange={()=>{setChecked(prevState=> !prevState)}}/>
                        <span className="pl-2">Accept Terms and Conditions</span>
                    </label>
                    <button disabled={!checked} className="ml-auto text-base bg-gray-100 enabled:bg-blue-600 enabled:text-white text-gray-400  font-bold px-1 py-0.5 rounded-md " onClick={()=>{setAcceptedTerms(true)}}>Next →</button>
                </div>
            </>
        ),[checked])

    const submitDogForm = useCallback( () => {
        return (
            <>
                <div className="mx-4">
                    <label><strong className="text-sm">Breed:</strong> <i className="text-xs">(We love mixed breeds too! Select &apos;Mix&apos;)</i></label>
                    <Select options={props.breedList} values={props.breedUploaded} searchable={!isMobile} dropdownHandle={false} dropdownPosition="top" placeholder="Select Breed (If a mix, select 'Mix')" loading={props.loading}  onChange={(values)=>{
                        props.setBreedUploaded(values)
                        props.setFilteredBreeds(values)}} />
                </div>
                <div className="flex justify-between mt-1 sm:mt-2">
                    <button className="text-base text-blue-600 border-2 font-bold px-1 py-0.5 rounded-md " onClick={()=>{setAcceptedTerms(false)}}>Back</button>
                    <button className="text-base bg-blue-600 text-white font-bold px-1 py-0.5 rounded-md " onClick={()=>{}}>Submit</button>
                </div>
            </>
        )
    },[props])

    const getBody = useCallback(()=>{
       if(!acceptedTerms){
           return TermsAndConditions()
       }
       return submitDogForm()
        }
       ,[TermsAndConditions, acceptedTerms, submitDogForm])


    return (
        <div className="shadow-xl rounded-lg max-w-2xl w-[95vw] glass-bg mb-2 p-2 sm:px-4 mx-auto pointer-events-auto">
            <h2 className="mx-auto text-xl sm:text-3xl text-center leading-tight mb-1 sm:mb-2">Add your pooch to the Internet&apos;s biggest collection of <strong>open source dog pictures!</strong></h2>
            {getBody()}
            <p className="mt-1 sm:mt-2 px-1 text-xs sm:text-sm tracking-tight leading-3 text-center"><strong>Note: </strong>Go Fetch! uses AI vision to prevent misuse, which may incorrectly prevent your upload. In this case, or if your dog&apos;s breed is not in the list, you can manually submit your photos as a <a className="text-blue-500 underline" href="https://github.com/jigsawpieces/dog-api-images#dog-api-images" target="_blank" rel="noreferrer" >GitHub pull request here.</a> </p>
        </div>
    )
}