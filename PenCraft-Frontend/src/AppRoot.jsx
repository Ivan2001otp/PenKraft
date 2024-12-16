import React from 'react'
import Footer from "./components/Footer";
import HeroImg from "./components/HeroImg";
import HeroSection from "./components/HeroSection";
import RecentBlogPage from "./components/RecentBlogPage";
import FeatureSection from './components/FeatureSection';
import Workflow from './components/Workflow';

const AppRoot = () => {
  return (
    <div className="mx-auto pt-20">
        <HeroImg />
        <FeatureSection/>
        <RecentBlogPage />
        <HeroSection/>
        <Workflow/>
        <Footer/>
      </div>
  )
}

export default AppRoot