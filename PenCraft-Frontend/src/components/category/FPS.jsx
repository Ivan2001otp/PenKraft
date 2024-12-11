import React from 'react'
import RecentPost from '../../Pages/RecentBlog'
import RecentPosts from '../../Pages/RecentBlogs';

const postData = [
    {
      title: "Dream destinations to visit this year in Paris",
      date: "03.08.2021",
      description: "Progressively incentivize cooperative systems through technically sound functionalities. Credibly productivate seamless data with flexible schemas.",
      image:
        "https://images.unsplash.com/photo-1732046827794-ac0d6c915a4a?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 2 Title",
      date: "2023-11-22",
      description: "This is a description for post 2.",
      image:
        "https://images.unsplash.com/photo-1489386659872-204f4f861691?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDE4fEZ6bzN6dU9ITjZ3fHxlbnwwfHx8fHw%3D",
    },
    {
      title: "Post 3 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732861448032-fc1b14365bff?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDd8NnNNVmpUTFNrZVF8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 4 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1730357753597-ee11ed5a2c0c?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDI5fENEd3V3WEpBYkV3fHxlbnwwfHx8fHw%3D",
    },
    {
      title: "Post 5 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732046827794-ac0d6c915a4a?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    },
    {
      title: "Post 6 Title",
      date: "2023-11-22",
      description: "This is a description for post 1.",
      image:
        "https://images.unsplash.com/photo-1732565432442-1d67155a5c84?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDExfHhqUFI0aGxrQkdBfHxlbnwwfHx8fHw%3D",
    },
    // ... more posts
  ];
const FPS = (blogList) => {
  return (
    <div className="p-6 rounded-lg">
    <h2 className="text-3xl font-semibold logo-font tracking-wider">FPS Section</h2>
    <RecentPosts posts={postData}/>
  </div>
  )
}

export default FPS