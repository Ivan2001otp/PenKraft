import { useState } from "react";
import { Link } from "react-router";
import Sony from "./category/Sony";
import FPS from "./category/FPS";
import RPG from "./category/RPG";
import Games from "./category/Games";

// const tabs = [
//   { id: 1, label: "News", path: "/" },
//   { id: 2, label: "Technology", path: "/Technology" },
//   { id: 3, label: "Cartoon", path: "/Cartoon" },
//   { id: 4, label: "Gaming", path: "/Gaming" },
// ];

const tabs = [
  {
    label: "RPGs",
    id: "rpg",
  },
  {
    label: "FPS",
    id: "fps",
  },
  {
    label: "Games",
    id: "games",
  },
  {
    label: "Sony",
    id: "sony",
  },
];

const TabBar = () => {
  const [selectedTab, setSelectedTab] = useState(tabs[0].id);
  // const [activeTabIndex, setActiveTabIndex] = useState(1);
  /*
  return (
    <div className="w-full mx-auto mt-4 px-4 -ml-4">
      <div className="flex text-2xl justify-between text-center logo-font">
        <div className="flex items-start space-x-4 overflow-x-auto pb-4 ">
          {tabs.map((tab) => (
          
            <div
              key={tab.id}
              className={`text-xl logo-font cursor-pointer text-center py-2 px-4 rounded-lg transition-all duration-300 ${
                activeTabIndex === tab.id
                  ? "text-slate-200 border-b-2 border-blue-600"
                  : "text-gray-500 hover:text-red-600 "
              }`}
              onClick={() => setActiveTabIndex(tab.id)}
            >
              {tab.label}
            </div>
          ))}
        </div>
        <div>
            <a 
             onClick={()=>{console.log("View All")}}
            
             className="hidden md:block hover:border-t-2 hover:border-r-2 hover:border-l-2 hover:border-emerald-200 px-4 py-1 duration-400 transition-all rounded-md">View All</a>
        </div>
      </div>
    </div>
  );
  */

  return (
    <div className="items-center">
      <div className="flex items-start space-x-4 mt-8">
        {tabs.map((item, index) => (
          <button
            key={index}
            onClick={() => setSelectedTab(item.id)}
            className={`px-4 py-2 text-lg font-semibold ${
              selectedTab === item.id
                ? "bg-slate-300 text-slate-800 rounded-3xl border-orange-400"
                : "bg-gray-200 text-slate-700"
            } rounded-lg transition duration-400`}
          >
            {item.label}
          </button>
        ))}
      </div>
      <div className="mt-8">
       
        {selectedTab === "sony" && <Sony />}
        {selectedTab === "fps" && <FPS />}
        {selectedTab === "rpg" && <RPG />}
        {selectedTab === "games" && <Games />}
      </div>
    </div>
  );
};

export default TabBar;
