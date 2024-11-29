import HeroImg from "./components/HeroImg";
import HeroSection from "./components/HeroSection";
import Navbar from "./components/Navbar";
import RecentBlogPage from "./components/RecentBlogPage";

function App() {
  return (
    <>
      
        <Navbar />
       
      <div className="mx-auto pt-20">
        {/* <HeroSection/> */}
        <HeroImg />
        <RecentBlogPage />
      </div>
    </>
  );
}

export default App;
