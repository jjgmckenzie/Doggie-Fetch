import DogDropDown from "@/app/dogdropdown";
import ManyDogs from "@/app/ManyDogs";

export default function Home() {
  return (
    <main className="max-w-6xl mx-auto px-2 pt-4">
        <div className="bg-white rounded-lg shadow p-3">
            <h1 className="text-4xl underline"> Hello, World!</h1>
        </div>
        <ManyDogs dogCount={50} class="scrollRight"/>
    </main>
  )
}
