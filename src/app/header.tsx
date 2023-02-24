export default function Header(){
    return (
        <div className="w-full shadow-lg bg-slate-50">
            <header className="max-w-7xl mx-auto flex px-2 ">
                <img src="/favicon.svg" className="h-20 w-20 pr-4" alt=""/>
                <h1 className="my-auto font-bold text-4xl drop-shadow-xl">Fetch!</h1>
                <a className="hover:underline mt-auto ml-auto" href="https://github.com/jjgmckenzie/Doggie-Fetch">github</a>
            </header>
        </div>
    )
}
