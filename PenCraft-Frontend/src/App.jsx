import AppRoot from "./AppRoot";
import Navbar from "./components/Navbar";
import ReadMorePage from "./Pages/ReadMorePage"
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";

function App() {
  return (
    <>
      <Navbar />

      <Routes>
          {/* Define routes here  */}
          <Route path="/" element={<AppRoot/>}/>
          <Route path="/read-more" element={<ReadMorePage/>}/>
      </Routes>
    </>
  );
}

export default App;
