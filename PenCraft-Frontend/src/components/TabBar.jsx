import Sony from "./category/Sony";
import FPS from "./category/FPS";
import RPG from "./category/RPG";
import Games from "./category/Games";
import { SONY_BLOGS_URL } from "../Util/Constants";
import { motion } from "framer-motion";
import { RPGs } from "../Util/Constants";
import { FPSG } from "../Util/Constants";
import { PS5G } from "../Util/Constants";
import { MARVEL } from "../Util/Constants";
import { DC } from "../Util/Constants";
import { SONY } from "../Util/Constants";

// React libraries
import { useEffect, useState } from "react";
import { CircularProgress } from "@mui/joy";
import { label } from "framer-motion/client";

const tabs = [
  {
    label: SONY,
    id: SONY,
  },
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
  
];

const TabBar = () => {
  const [selectedTab, setSelectedTab] = useState(tabs[0].id);
  const [jsonResponse, setJsonResponse] = useState(Array)

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

  const params = new URLSearchParams({
    category: "Sony",
    cursor: "",
  });

  const urlWithParams = `${SONY_BLOGS_URL}?${params.toString()}`;

  // fetches data from api
  useEffect(() => {
    const loadData = async () => {
      setLoading(true);
      await delay(2000);

      try {
        await fetch(urlWithParams)
          .then((response) => response.json())
          .then((resp) => {
            let saveResp = resp["data"];
            setJsonResponse(saveResp)
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
  }, [selectedTab]);


  if (loading) {
    return (
      <motion.div
        initial={{ opacity: 0 }}
        exit={{ opacity: 1 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 1 }}
        className="flex items-center justify-center w-full h-screen"
      >
        <CircularProgress color="success" variant="outlined" size="lg" />
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
                  ? "bg-blue-600 text-slate-300 rounded-3xl border-2  border-white"
                  : "bg-gray-200 text-slate-700"
              } rounded-lg transition duration-400`}
            >
              {item.label}
            </button>
          ))}
        </div>
        <div className="mt-8">
          {selectedTab === RPGs && <RPG blogList={[]} />}
          {selectedTab === FPSG && <FPS blogList={[]} />}
          {selectedTab === PS5G && <Games blogList={[]} />}
          {selectedTab === SONY && <Sony blogList={jsonResponse || []} />}
        </div>
      </div>
    </motion.div>
  );
};

export default TabBar;
