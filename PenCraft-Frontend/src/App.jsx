import Footer from "./components/Footer";
import HeroImg from "./components/HeroImg";
import HeroSection from "./components/HeroSection";
import Navbar from "./components/Navbar";
import RecentBlogPage from "./components/RecentBlogPage";
import TabBar from "./components/TabBar";
import {  Router, Route, Routes } from "react-router-dom";

function App() {
  return (
    <>
      <Navbar />

      <div className="mx-auto pt-20">
        <HeroImg />
        <RecentBlogPage />
        <HeroSection/>
        <Footer/>
      </div>
    </>
  );
}

export default App;
