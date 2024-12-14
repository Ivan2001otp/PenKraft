import React, { useEffect, useState, useRef } from 'react'
import RecentPost from './RecentBlog';
import { useLocation } from 'react-router'

const ReadMorePage = () => {
  const location = useLocation();
  const {list} = location.state || {};

  const [blogs, setBlogs] = useState(list || []);
  const [cursor, setCursor] = useState('');
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);

  const observeRef = useRef(null);

  const fetchMoreBlogs = async() => {

  }


  // Use of IntersectionObserver  to trigger fetching
  useEffect(() => {
    if (observeRef.current) observeRef.current.disconnect;

    const observer = new IntersectionObserver(
      ([entry])=>{
        if (entry.isIntersecting) {
          fetchMoreBlogs();
        }
      },
      {
        threshold: 1.0
      }
    );

    if (observeRef.current) {
      observer.observe(observeRef.current) // observe the last item
    }

    observeRef.current = observer;
    return () => observer.disconnect();
  },[blogs])

  return (
    <div>
        <div className='px-6 py-2 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-10 mt-8'>
          { 
            list &&
            list.map((blog, index) => (
              <RecentPost key={index} {...blog}/>
            ))
          }
        </div>
        
        {/* Target element for triggering fetch */}
        <div
          ref={(el) => {
            observeRef.current = el;
          }}

          style={{height:20}}
        >
          {loading && <p className='text-white flex items-center justify-center'>Loading more blogs</p>}
          {!hasMore && <p className='text-white flex items-center justify-center'>No more blogs to display</p>}
        </div>
  </div>
  )
}

export default ReadMorePage