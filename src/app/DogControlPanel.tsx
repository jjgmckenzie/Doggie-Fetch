`use client`
import {Dispatch, SetStateAction, useCallback} from "react";


interface Props{
    optionsPoppedUp:boolean
    uploadPoppedUp:boolean
    setOptionsPoppedUp:Dispatch<SetStateAction<boolean>>
    setUploadPoppedUp:Dispatch<SetStateAction<boolean>>
}

enum panelSelected{
    options,
    upload
}


export default function DogControlPanel (props:Props) {

    const togglePanelShown = useCallback((toggledBy:panelSelected)=>{
        if(props.optionsPoppedUp && toggledBy == panelSelected.options){
            props.setOptionsPoppedUp(false)
        }
        else if(props.uploadPoppedUp && toggledBy == panelSelected.upload){
            props.setUploadPoppedUp(false)
        }
        else{
            if(toggledBy == panelSelected.upload){
                props.setOptionsPoppedUp(false)
                props.setUploadPoppedUp(true)
            }
            if(toggledBy == panelSelected.options){
                props.setOptionsPoppedUp(true)
                props.setUploadPoppedUp(false)
            }
        }
    },[props])

    return (
        <div className="w-16 mr-4 pointer-events-auto">
            <button className="ctrl-panel px-1 min-w-fit" onClick={()=>{togglePanelShown(panelSelected.options)}}>Options</button><br/>
            <button className="ctrl-panel mt-3 px-1 min-w-fit" onClick={()=>{togglePanelShown(panelSelected.upload)}}>Upload</button>
        </div>
    )
}