import Sony from "./category/Sony";
import FPS from "./category/FPS";
import RPG from "./category/RPG";
import Games from "./category/Games";
import { FETCH_ALL_BLOG_URL } from "../Util/Constants";
import { motion } from "framer-motion";
import {RPGs} from '../Util/Constants'
import {FPSG} from '../Util/Constants'
import {PS5G} from '../Util/Constants'
import {MARVEL} from '../Util/Constants'
import {DC} from '../Util/Constants'


// React libraries
import { useEffect, useState } from "react";
import { CircularProgress } from "@mui/joy";

const tabs = [
  {
    label: RPGs,
    id: RPGs,
  },
  {
    label: FPSG,
    id: FPSG,
  },
  {
    label: PS5G,
    id: PS5G,
  },
  
  {
    label: MARVEL,
    id:MARVEL,
  },
  {
    label: DC,
    id: DC,
  }
];


const TabBar = () => {
  const [selectedTab, setSelectedTab] = useState(tabs[0].id);
  let jsonResponse;
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
 
  // fetches data from api
  useEffect(() => {
    const loadData = async () => {

      setLoading(true);
      await delay(3000);

      try {
        await fetch(FETCH_ALL_BLOG_URL)
          .then((response) => response.json())
          .then((resp) => {
            let saveResp = resp["data"];
            jsonResponse = saveResp;
            console.log("The response is saveResp ", jsonResponse);
          });
      } catch (error) {
        console.log(error);
        setError("Failed to fetch data from api !");
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, []);

  // filters the data on the basis of types
  useEffect(()=>{
    setLoading(true);
    delay(4000)
    console.log(`Displaying ${selectedTab} section !`);
    setLoading(false);
    // filter logic and give the output.
  },[selectedTab])


  if (loading) {
    return (
      <motion.div
        initial={{ opacity: 0 }}
        exit={{ opacity: 1 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 1 }}
        className="loading-container"
      >
        <CircularProgress color="success" variant="outlined" size="lg"/>
       </motion.div>
    );
  }

  if (error) {
    return (
      <motion.div
        initial={{ opacity: 0 }}
        exit={{ opacity: 1 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 2 }}
        className="error-container"
      >
        <div className="items-center justify-center text-3xl text-red-500 p-4">
          {error}
        </div>
      </motion.div>
    );
  }

  return (
    <motion.div
      initial={{ opacity: 0 }}
      exit={{ opacity: 1 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 3 }}
      className="data-container"
    >
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
          {selectedTab === RPGs && <Sony blogList={[]} />}
          {selectedTab === FPSG && <FPS blogList={[]}/>}
          {selectedTab === PS5G && <RPG blogList={[]}/>}
          {selectedTab === MARVEL && <Games blogList={[]}/>}
          {selectedTab === DC && <Games blogList={[]}/>}

        </div>
      </div>
    </motion.div>
  );
};

export default TabBar;
