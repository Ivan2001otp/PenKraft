import React from "react";
import Footer from "./components/Footer";
import HeroImg from "./components/HeroImg";
import HeroSection from "./components/HeroSection";
import RecentBlogPage from "./components/RecentBlogPage";
import FeatureSection from "./components/FeatureSection";
import Workflow from "./components/Workflow";
import FloatingActionButton from "./components/FloatingActionButton";
import { Link } from "react-router";

const AppRoot = () => {
  return (
    <div>
      <div className="mx-auto pt-20">
        <HeroImg />
        <FeatureSection />
        <RecentBlogPage />
        <HeroSection />
        <Workflow />
        <Footer />
      </div>

      <Link to={"/workspace"}>
        <button className="fab hover:text-yellow-500  hover:uppercase px-4 logo-font py-3 rounded-3xl backdrop-blur-sm ">
          Kraft Blogs
        </button>
      </Link>
    </div>
  );
};

export default AppRoot;
