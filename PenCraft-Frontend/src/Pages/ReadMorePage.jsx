import React, { useEffect, useState, useRef } from 'react'
import RecentPost from './RecentBlog';
import { useLocation } from 'react-router'
import { SONY_BLOGS_URL } from "../Util/Constants";
import { motion } from 'framer-motion'; 

const ReadMorePage = () => {
  const location = useLocation();
  const {list, niche} = location.state || {};

  let cursor1 = list[list.length -1 ]
 
  const [blogs, setBlogs] = useState(list || []); // to store the list of blogs
  const [loading, setLoading] = useState(false); // to track loading state
  const [cursor, setCursor] = useState(cursor1['blog_id'].toString()); // page number for pagination
  

  const fetchMoreBlogs = async() => {
      setLoading(true)

      const params = new URLSearchParams({
        category: niche,
        cursor: cursor,
      });


      const urlWithParams = `${SONY_BLOGS_URL}?${params.toString()}`;

      try {

        await fetch(urlWithParams)
          .then((response) => response.json())
          .then((result) => {
            console.log("The response is saveResp ", result);
            setCursor(result['cursor'])

            const newData = result['data'];
            setBlogs((prevList)=>[...prevList, ...newData]);
          });

      } catch (error) {
        console.log(error);
        
      } finally {
        setLoading(false);
      }
  }

  return (
    <div>
        <div className='px-6 py-2 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-10 mt-8'>
          { 
            blogs &&
            blogs.map((blog, index) => (
              <motion.div
                key={index}
                initial={{ opacity:0, x:-100}}
                animate={{opacity:1, x:0}}
                exit={{opacity:0, x:-100}}
                transition={{
                  duration:0.6,
                  delay: 0.3,
                  ease: 'easeOut'
                }}
              >
              <RecentPost key={index} {...blog}/>
              </motion.div>
              
            ))
          }
        </div>
        
        <div className='flex text-slate-200 text-xl  justify-center mt-8'>
            <button
            onClick={fetchMoreBlogs}
            className='px-6 py-2 bg-blue-300 text-white hover:bg-blue-600 rounded-full'
            disabled={loading}
            >
              {loading ? 'Loading...' : 'More Blogs'}
            </button>
        </div>

            {/* loading spinner */}
        {
          loading && (
            <div className='flex justify-center mt-8'>
            <div className="w-10 h-10 border-t-2 border-red-500 border-solid rounded-full animate-spin"></div>
       
            </div>
          )
        }
        
  </div>
  )
}

export default ReadMorePage