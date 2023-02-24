export default function Header(){
    return (
        <div className="w-full shadow-lg glass-bg rounded-b-md">
            <header className="max-w-7xl mx-auto flex px-2 ">
                <img src="/favicon.svg" className="h-16 w-16 pr-4" alt=""/>
                <h1 className="my-auto font-bold text-2xl sm:text-4xl drop-shadow-xl">Go Fetch!</h1>
                <div className="w-fit ml-auto h-full text-right my-auto text-sm">
                    <a className="hover:underline" href="https://dog.ceo/dog-api/about">About Dog.CEO API</a><br/>
                    <a className="hover:underline" href="https://github.com/jjgmckenzie/Doggie-Fetch">Fork Go Fetch! on GitHub</a><br/>
                    <a className="hover:underline" href="https://jjgmckenzie.ca/">Go Fetch! by jjgmckenzie</a>
                </div>
            </header>
        </div>
    )
}
