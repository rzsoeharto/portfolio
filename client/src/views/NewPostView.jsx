import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import Confirmation from "../components/Confimation";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import SectionSelection from "../components/SectionSelected";
import AddSection from "../components/AddSection";
import useLogin, { modalStorage } from "../stores/useStore";
import { showToast, toastWithoutFade } from "../utils/toastUtils";

function NewPostView() {
  const navigate = useNavigate();
  const { isLoggedIn } = useLogin();
  const [titleData, setTitleData] = useState("");
  const [sectionsData, setSectionsData] = useState([]);
  const { modalState, setModalState, setModalType } = modalStorage();
  const [sectionModal, setSectionModal] = useState(false);

  function CancelPost() {
    if (titleData == "" && sectionsData == 0) {
      navigate("/posts");
      return;
    }
    setModalType("");
    setModalState(true);
  }

  function onChange(e) {
    setTitleData(e.target.value);
  }

  function handleDelete(id) {
    setSectionsData((state) => {
      const updatedSection = [...state];
      updatedSection.splice(id, 1);
      return updatedSection;
    });
  }

  function submitPost() {
    if (!isLoggedIn) {
      showToast("You are not logged in", "Warning");
      return;
    }

    if (titleData == "") {
      showToast("Title can not be empty", "Warning");
      return;
    } else if (sectionsData == 0) {
      showToast("Content cannot be empty", "Warning");
      return;
    }

    let id = toastWithoutFade("Saving", "Loading");

    fetch("http://localhost:8080/create-post", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        Title: titleData,
        Sections: sectionsData,
      }),
    })
      .then((res) => {
        console.log(res);
        if (!res.ok) {
          showToast("Something went wrong in the backend", "Warning");
          return;
        }
        id.className = id.className.replace("block ", "hidden show");
      })
      .catch((error) => {
        console.error(error);
        showToast("Unable to connect to server", "Warning");
        return;
      });
  }

  return (
    <>
      {modalState ? <Confirmation /> : <></>}
      <div className="flex flex-row">
        <Navbar />
        <div className="flex flex-col h-screen w-full">
          <Logo />
          <div className="flex flex-col w-full gap-3 pb-10">
            <div className="flex flex-row gap-5">
              <input
                id="Title"
                onChange={onChange}
                type="text"
                placeholder="Title"
                className="text-3xl text-[#303030] w-[740px] font-semibold focus:outline-none"
              />
              <button
                onClick={CancelPost}
                className="text-2xl duration-200 hover:text-[#FFA360] bg-transparent"
              >
                Cancel
              </button>
            </div>
            <div className="flex flex-col w-[740px] gap-2">
              <div className="flex flex-col gap-3">
                {Array.isArray(sectionsData) &&
                  sectionsData.map((section, index) => (
                    <div key={index} className="text-lg flex flex-row">
                      <SectionSelection
                        index={index}
                        sectionSelection={section}
                        setSectionsData={setSectionsData}
                      />
                      <button
                        onClick={() => {
                          handleDelete(index);
                        }}
                        className="font-bold bg-transparent h-10 w-14"
                      >
                        X
                      </button>
                    </div>
                  ))}
              </div>
              <div className="flex flex-row gap">
                {sectionModal ? (
                  <AddSection
                    setSectionModal={setSectionModal}
                    setSectionsData={setSectionsData}
                  />
                ) : (
                  <button
                    onClick={() => {
                      setSectionModal(true);
                    }}
                    className="bg-[#d9d9d9] h-[60px] text-xl w-full rounded hover:bg-white duration-200"
                  >
                    + Add Section
                  </button>
                )}
              </div>
              <button
                className="h-[60px] bg-[#FFA360] text-xl font-semibold hover:bg-white duration-200"
                onClick={submitPost}
              >
                Submit
              </button>
              <button className="" onClick={() => console.log(sectionsData)}>
                log the thing
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default NewPostView;
