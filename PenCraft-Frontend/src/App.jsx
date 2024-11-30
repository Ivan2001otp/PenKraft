import HeroImg from "./components/HeroImg";
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
        
      </div>
    </>
  );
}

export default App;
