import { create } from "zustand";

export const toastStorage = create((set) => ({
  toastType: 0,
  setToastType: (val) => set(() => ({ toastType: val })),

  toastMessage: "",
  setToastMessage: (val) => set(() => ({ toastMessage: val })),
}));

const useLogin = create((set) => ({
  isLoggedIn: false,
  setLoggedIn: (val) => set(() => ({ isLoggedIn: val })),
}));

export const userInfo = create(() => ({
  username: localStorage.getItem("username"),
  name: localStorage.getItem("name"),
}));

export const modalStorage = create((set) => ({
  modalState: false,
  modalType: "",
  setModalState: (val) => set(() => ({ modalState: val })),
  setModalType: (val) => set(() => ({ modalType: val })),
}));

export const postStorage = create((set) => ({
  titleData: "",
  sectionsData: [],

  setTitle: (val) => set(() => ({ titleData: val })),

  setSections: (val) =>
    set((state) => ({ sectionsData: [...state.sectionsData, val] })),

  updateSectionContent: (index, val) =>
    set((state) => {
      const updatedSections = [...state.sectionsData];
      updatedSections[index] = { ...updatedSections[index], Content: val };
      return { sectionsData: updatedSections };
    }),

  removeSection: (index) =>
    set((state) => {
      const updatedSection = [...state.sectionsData];
      updatedSection.splice(index, 1);
      return { sectionsData: updatedSection };
    }),

  clearSections: () =>
    set((state) => {
      const arr = [...state.sectionsData];
      arr.splice(0, arr.length);
      return { sectionsData: arr };
    }),
}));

export default useLogin;
