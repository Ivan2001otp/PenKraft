import React from "react";
import RecentPosts from "../Pages/RecentBlogs";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  useNavigate,
  Outlet,
} from "react-router-dom";

import TabBar from "./TabBar";

const RecentBlogPage = () => {
  const postsData = [
    {
      title: "Dream destinations to visit this year in Paris",
      date: "03.08.2021",
      description: "Progressively incentivize cooperative systems through technically sound functionalities. Credibly productivate seamless data with flexible schemas.",
      image:
        "https://images.unsplash.com/photo-1731963914155-d22942204d3d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwzNXx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 2 Title",
      date: "2023-11-22",
      description: "This is a description for post 2.",
      image:
        "https://images.unsplash.com/photo-1732719632991-6ec4cdb32dfb?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwyMnx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 3 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732719632991-6ec4cdb32dfb?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwyMnx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 4 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732719632991-6ec4cdb32dfb?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwyMnx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 5 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732719632991-6ec4cdb32dfb?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwyMnx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 6 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732719632991-6ec4cdb32dfb?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwyMnx8fGVufDB8fHx8fA%3D%3D",
    },
    // ... more posts
  ];

  return (
    <div className="p-2 mt-12 mx-6 h-[650px]">
      <div className="text-xl logo-font md:text-3xl">Recent Game Blogs</div>
      <div className="flex flex-col">
        <Router>
          <div>
            <TabBar />
            <Routes>
              <Route path="/" element={<RecentPosts posts={postsData} />} />

              <Route path="/News" element={<RecentPosts posts={postsData} />} />

              <Route
                path="/Technology"
                element={<RecentPosts posts={postsData} />}
              />

              <Route
                path="/Cartoon"
                element={<RecentPosts posts={postsData} />}
              />

              <Route
                path="/Gaming"
                element={<RecentPosts posts={postsData} />}
              />
            </Routes>
          </div>
        </Router>
      </div>
      {/* <RecentPosts posts={postsData} /> */}
    </div>
  );
};

export default RecentBlogPage;
