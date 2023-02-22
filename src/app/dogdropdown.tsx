interface Props{
    breeds:String[];
}
export default function DogDropDown(props:Props){
    const breedsFormatted = props.breeds.join(", ")
    return (
        <>
            <input/>
            <div>
                {breedsFormatted}
            </div>
        </>
    )
}