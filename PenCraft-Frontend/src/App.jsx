import HeroImg from './components/HeroImg'
import HeroSection from './components/HeroSection'
import Navbar from './components/Navbar'

function App() {
  return (
    <>
      <Navbar/>
      <HeroImg/>
      <div className='mx-auto max-w-7xl px-6 pt-20'>
        <HeroSection/>
      </div>
    </>
  )  
}

export default App
