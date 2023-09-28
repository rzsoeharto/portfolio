import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";
import PostView from "./views/PostView";
import HomeView from "./views/Home";
import ContactView from "./views/ContactView";
import PortfolioView from "./views/PortfolioView";
import SpecificPostView from "./views/SpecificPostView";
import LoginView from "./views/LoginView";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<HomeView />} />
          <Route path="/posts" element={<PostView />} />
          <Route path="/posts/:postID" element={<SpecificPostView />} />
          <Route path="/contact" element={<ContactView />} />
          <Route path="/portfolio" element={<PortfolioView />} />
          <Route path="/login" element={<LoginView />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
