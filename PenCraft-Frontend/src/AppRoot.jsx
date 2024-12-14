import React from 'react'
import Footer from "./components/Footer";
import HeroImg from "./components/HeroImg";
import HeroSection from "./components/HeroSection";
import RecentBlogPage from "./components/RecentBlogPage";

const AppRoot = () => {
  return (
    <div className="mx-auto pt-20">
        <HeroImg />
        <RecentBlogPage />
        <HeroSection/>
        <Footer/>
      </div>
  )
}

export default AppRoot