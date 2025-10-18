import { createContext, useContext, useState, ReactNode } from "react";

export type ScheduleTime = {
  id: string;
  day: string;
  startTime: string;
  endTime: string;
};

export type Course = {
  id: string;
  name: string;
  code: string;
  instructor: string;
  students: number;
  schedule: string; // Legacy field for display
  scheduleDetails?: ScheduleTime[]; // New structured schedule
  location?: string;
  color: string;
  progress: number;
  description?: string;
  credits?: number;
  semester?: string;
  gpa?: number;
};

export type Note = {
  id: string;
  title: string;
  courseId: string;
  courseCode: string;
  date: string;
  pages: number;
  summarized: boolean;
  tags: string[];
  content?: string;
};

export type CourseMaterial = {
  id: string;
  courseId: string;
  title: string;
  type: 'video' | 'pdf' | 'text' | 'link' | 'other';
  url?: string;
  uploadDate: string;
  size?: string;
  description?: string;
};

export type Tag = {
  id: string;
  name: string;
  type: 'default' | 'custom';
  color: string;
};

export type KnowledgeConnection = {
  from: Note;
  to: Note;
  sharedTags: string[];
  strength: number;
};

export type Semester = {
  id: string;
  name: string;
  startDate: string;
  endDate: string;
};

export type Reminder = {
  id: string;
  title: string;
  description?: string;
  dueDate: string;
  dueTime: string;
  courseId?: string;
  courseCode?: string;
  priority: 'low' | 'medium' | 'high';
  completed: boolean;
  type: 'assignment' | 'exam' | 'reading' | 'project' | 'other';
};

export type GraduationRequirements = {
  requiredCredits: number;
  currentCredits: number;
};

export type UserPreferences = {
  personalInfo: {
    firstName: string;
    lastName: string;
    email: string;
    phoneNumber?: string;
    studentId?: string;
    major?: string;
    profilePicture?: string;
  };
  academic: {
    gpaScale: '4.0' | '5.0' | '10.0';
  };
  appearance: {
    theme: 'light' | 'dark' | 'system';
  };
};

interface DataContextType {
  courses: Course[];
  notes: Note[];
  tags: Tag[];
  materials: CourseMaterial[];
  semesters: Semester[];
  reminders: Reminder[];
  graduationRequirements: GraduationRequirements;
  userPreferences: UserPreferences;
  updateGraduationRequirements: (requirements: Partial<GraduationRequirements>) => void;
  updateUserPreferences: (preferences: Partial<UserPreferences>) => void;
  addCourse: (course: Omit<Course, 'id'>) => void;
  updateCourse: (id: string, course: Partial<Course>) => void;
  deleteCourse: (id: string) => void;
  addNote: (note: Omit<Note, 'id'>) => void;
  updateNote: (id: string, note: Partial<Note>) => void;
  deleteNote: (id: string) => void;
  addTag: (tag: Omit<Tag, 'id'>) => void;
  addMaterial: (material: Omit<CourseMaterial, 'id'>) => void;
  updateMaterial: (id: string, material: Partial<CourseMaterial>) => void;
  deleteMaterial: (id: string) => void;
  addSemester: (semester: Omit<Semester, 'id'>) => void;
  updateSemester: (id: string, semester: Partial<Semester>) => void;
  deleteSemester: (id: string) => void;
  addReminder: (reminder: Omit<Reminder, 'id'>) => void;
  updateReminder: (id: string, reminder: Partial<Reminder>) => void;
  deleteReminder: (id: string) => void;
  getCourseById: (id: string) => Course | undefined;
  getNotesByCourse: (courseId: string) => Note[];
  getMaterialsByCourse: (courseId: string) => CourseMaterial[];
  getRemindersByCourse: (courseId: string) => Reminder[];
}

const DataContext = createContext<DataContextType | undefined>(undefined);

// Default tags that come with the system
const defaultTags: Tag[] = [
  { id: "1", name: "Algorithms", type: "default", color: "bg-blue-500" },
  { id: "2", name: "Data Structures", type: "default", color: "bg-green-500" },
  { id: "3", name: "Mathematics", type: "default", color: "bg-purple-500" },
  { id: "4", name: "Physics", type: "default", color: "bg-orange-500" },
  { id: "5", name: "Theory", type: "default", color: "bg-pink-500" },
  { id: "6", name: "Practice", type: "default", color: "bg-cyan-500" },
  { id: "7", name: "Exam Prep", type: "default", color: "bg-red-500" },
  { id: "8", name: "Assignment", type: "default", color: "bg-yellow-500" },
  { id: "9", name: "Project", type: "default", color: "bg-indigo-500" },
  { id: "10", name: "Review", type: "default", color: "bg-teal-500" },
];

// Initial mock data
const initialSemesters: Semester[] = [
  {
    id: "1",
    name: "Fall 2024",
    startDate: "2024-09-01",
    endDate: "2024-12-15"
  },
  {
    id: "2",
    name: "Spring 2025",
    startDate: "2025-01-15",
    endDate: "2025-05-15"
  }
];

const initialCourses: Course[] = [
  {
    id: "1",
    name: "Computer Science 101",
    code: "CS101",
    instructor: "Dr. Sarah Johnson",
    students: 45,
    schedule: "Mon, Wed 9:00 AM",
    scheduleDetails: [
      { id: "1", day: "Monday", startTime: "09:00 AM", endTime: "10:00 AM" },
      { id: "2", day: "Wednesday", startTime: "09:00 AM", endTime: "10:00 AM" }
    ],
    location: "Room 101",
    color: "bg-blue-500",
    progress: 65,
    description: "Introduction to computer science fundamentals",
    credits: 4,
    semester: "Fall 2024",
    gpa: 3.5
  },
  {
    id: "2",
    name: "Advanced Mathematics",
    code: "MATH201",
    instructor: "Prof. Michael Chen",
    students: 32,
    schedule: "Tue, Thu 2:00 PM",
    scheduleDetails: [
      { id: "1", day: "Tuesday", startTime: "14:00 PM", endTime: "15:00 PM" },
      { id: "2", day: "Thursday", startTime: "14:00 PM", endTime: "15:00 PM" }
    ],
    location: "Room 201",
    color: "bg-purple-500",
    progress: 78,
    description: "Advanced calculus and mathematical analysis",
    credits: 3,
    semester: "Fall 2024",
    gpa: 3.8
  },
  {
    id: "3",
    name: "Data Structures",
    code: "CS202",
    instructor: "Dr. Emily Brown",
    students: 38,
    schedule: "Mon, Wed 2:00 PM",
    scheduleDetails: [
      { id: "1", day: "Monday", startTime: "14:00 PM", endTime: "15:00 PM" },
      { id: "2", day: "Wednesday", startTime: "14:00 PM", endTime: "15:00 PM" }
    ],
    location: "Room 102",
    color: "bg-green-500",
    progress: 42,
    description: "Advanced data structures and algorithms",
    credits: 4,
    semester: "Fall 2024",
    gpa: 3.2
  },
  {
    id: "4",
    name: "Physics II",
    code: "PHY102",
    instructor: "Dr. Robert Wilson",
    students: 50,
    schedule: "Tue, Thu 10:00 AM",
    scheduleDetails: [
      { id: "1", day: "Tuesday", startTime: "10:00 AM", endTime: "11:00 AM" },
      { id: "2", day: "Thursday", startTime: "10:00 AM", endTime: "11:00 AM" }
    ],
    location: "Room 301",
    color: "bg-orange-500",
    progress: 55,
    description: "Mechanics and thermodynamics",
    credits: 3,
    semester: "Fall 2024",
    gpa: 3.7
  }
];

const initialNotes: Note[] = [
  {
    id: "1",
    title: "Introduction to Algorithms",
    courseId: "1",
    courseCode: "CS101",
    date: "Oct 10, 2025",
    pages: 12,
    summarized: true,
    tags: ["Algorithms", "Theory"]
  },
  {
    id: "2",
    title: "Calculus Derivatives",
    courseId: "2",
    courseCode: "MATH201",
    date: "Oct 11, 2025",
    pages: 8,
    summarized: true,
    tags: ["Mathematics", "Theory"]
  },
  {
    id: "3",
    title: "Binary Trees and Graphs",
    courseId: "3",
    courseCode: "CS202",
    date: "Oct 12, 2025",
    pages: 15,
    summarized: false,
    tags: ["Data Structures", "Algorithms"]
  },
  {
    id: "4",
    title: "Newton's Laws",
    courseId: "4",
    courseCode: "PHY102",
    date: "Oct 12, 2025",
    pages: 10,
    summarized: true,
    tags: ["Physics", "Theory"]
  },
  {
    id: "5",
    title: "Sorting Algorithms",
    courseId: "1",
    courseCode: "CS101",
    date: "Oct 13, 2025",
    pages: 9,
    summarized: false,
    tags: ["Algorithms", "Practice"]
  }
];

const initialMaterials: CourseMaterial[] = [
  {
    id: "1",
    courseId: "1",
    title: "Introduction to Programming - Lecture Slides",
    type: "pdf",
    uploadDate: "Oct 8, 2025",
    size: "2.4 MB",
    description: "Week 1 lecture slides"
  },
  {
    id: "2",
    courseId: "1",
    title: "Python Basics Tutorial",
    type: "video",
    url: "https://example.com/video1",
    uploadDate: "Oct 9, 2025",
    description: "Introductory video tutorial"
  },
  {
    id: "3",
    courseId: "2",
    title: "Calculus Textbook Chapter 3",
    type: "pdf",
    uploadDate: "Oct 10, 2025",
    size: "5.8 MB"
  }
];

const initialReminders: Reminder[] = [
  {
    id: "1",
    title: "CS101 Assignment Due",
    description: "Complete programming assignment 3",
    dueDate: "2024-10-25",
    dueTime: "23:59",
    courseId: "1",
    courseCode: "CS101",
    priority: "high",
    completed: false,
    type: "assignment"
  },
  {
    id: "2",
    title: "MATH201 Midterm Exam",
    description: "Study chapters 1-5",
    dueDate: "2024-10-28",
    dueTime: "14:00",
    courseId: "2",
    courseCode: "MATH201",
    priority: "high",
    completed: false,
    type: "exam"
  }
];

export function DataProvider({ children }: { children: ReactNode }) {
  const [courses, setCourses] = useState<Course[]>(initialCourses);
  const [notes, setNotes] = useState<Note[]>(initialNotes);
  const [tags, setTags] = useState<Tag[]>(defaultTags);
  const [materials, setMaterials] = useState<CourseMaterial[]>(initialMaterials);
  const [semesters, setSemesters] = useState<Semester[]>(initialSemesters);
  const [reminders, setReminders] = useState<Reminder[]>(initialReminders);
  const [graduationRequirements, setGraduationRequirements] = useState<GraduationRequirements>({
    requiredCredits: 120,
    currentCredits: 0
  });
  const [userPreferences, setUserPreferences] = useState<UserPreferences>({
    personalInfo: {
      firstName: "Anh Son",
      lastName: "Nguyen",
      email: "nguyenanhson@gmail.com",
      studentId: "2023001",
      major: "Computer Science"
    },
    academic: {
      gpaScale: '4.0'
    },
    appearance: {
      theme: 'system'
    }
  });

  const addCourse = (course: Omit<Course, 'id'>) => {
    const newCourse = {
      ...course,
      id: Date.now().toString(),
    };
    setCourses(prev => [...prev, newCourse]);
  };

  const updateCourse = (id: string, updatedCourse: Partial<Course>) => {
    setCourses(prev =>
      prev.map(course => (course.id === id ? { ...course, ...updatedCourse } : course))
    );
  };

  const deleteCourse = (id: string) => {
    setCourses(prev => prev.filter(course => course.id !== id));
  };

  const addNote = (note: Omit<Note, 'id'>) => {
    const newNote = {
      ...note,
      id: Date.now().toString(),
    };
    setNotes(prev => [...prev, newNote]);
  };

  const updateNote = (id: string, updatedNote: Partial<Note>) => {
    setNotes(prev =>
      prev.map(note => (note.id === id ? { ...note, ...updatedNote } : note))
    );
  };

  const deleteNote = (id: string) => {
    setNotes(prev => prev.filter(note => note.id !== id));
  };

  const addTag = (tag: Omit<Tag, 'id'>) => {
    const newTag = {
      ...tag,
      id: Date.now().toString(),
    };
    setTags(prev => [...prev, newTag]);
  };

  const getCourseById = (id: string) => {
    return courses.find(course => course.id === id);
  };

  const getNotesByCourse = (courseId: string) => {
    return notes.filter(note => note.courseId === courseId);
  };

  const addMaterial = (material: Omit<CourseMaterial, 'id'>) => {
    const newMaterial = {
      ...material,
      id: Date.now().toString(),
    };
    setMaterials(prev => [...prev, newMaterial]);
  };

  const updateMaterial = (id: string, updatedMaterial: Partial<CourseMaterial>) => {
    setMaterials(prev =>
      prev.map(material => (material.id === id ? { ...material, ...updatedMaterial } : material))
    );
  };

  const deleteMaterial = (id: string) => {
    setMaterials(prev => prev.filter(material => material.id !== id));
  };

  const getMaterialsByCourse = (courseId: string) => {
    return materials.filter(material => material.courseId === courseId);
  };

  const addSemester = (semester: Omit<Semester, 'id'>) => {
    const newSemester = {
      ...semester,
      id: Date.now().toString(),
    };
    setSemesters(prev => [...prev, newSemester]);
  };

  const updateSemester = (id: string, updatedSemester: Partial<Semester>) => {
    setSemesters(prev =>
      prev.map(semester => (semester.id === id ? { ...semester, ...updatedSemester } : semester))
    );
  };

  const deleteSemester = (id: string) => {
    setSemesters(prev => prev.filter(semester => semester.id !== id));
  };

  const addReminder = (reminder: Omit<Reminder, 'id'>) => {
    const newReminder = {
      ...reminder,
      id: Date.now().toString(),
    };
    setReminders(prev => [...prev, newReminder]);
  };

  const updateReminder = (id: string, updatedReminder: Partial<Reminder>) => {
    setReminders(prev =>
      prev.map(reminder => (reminder.id === id ? { ...reminder, ...updatedReminder } : reminder))
    );
  };

  const deleteReminder = (id: string) => {
    setReminders(prev => prev.filter(reminder => reminder.id !== id));
  };

  const getRemindersByCourse = (courseId: string) => {
    return reminders.filter(reminder => reminder.courseId === courseId);
  };

  const updateGraduationRequirements = (requirements: Partial<GraduationRequirements>) => {
    setGraduationRequirements(prev => ({ ...prev, ...requirements }));
  };

  const updateUserPreferences = (preferences: Partial<UserPreferences>) => {
    setUserPreferences(prev => ({
      ...prev,
      ...preferences,
      personalInfo: preferences.personalInfo ? { ...prev.personalInfo, ...preferences.personalInfo } : prev.personalInfo,
      academic: preferences.academic ? { ...prev.academic, ...preferences.academic } : prev.academic,
      appearance: preferences.appearance ? { ...prev.appearance, ...preferences.appearance } : prev.appearance
    }));
  };

  return (
    <DataContext.Provider
      value={{
        courses,
        notes,
        tags,
        materials,
        semesters,
        reminders,
        graduationRequirements,
        userPreferences,
        addCourse,
        updateCourse,
        deleteCourse,
        addNote,
        updateNote,
        deleteNote,
        addTag,
        addMaterial,
        updateMaterial,
        deleteMaterial,
        addSemester,
        updateSemester,
        deleteSemester,
        addReminder,
        updateReminder,
        deleteReminder,
        updateGraduationRequirements,
        updateUserPreferences,
        getCourseById,
        getNotesByCourse,
        getMaterialsByCourse,
        getRemindersByCourse,
      }}
    >
      {children}
    </DataContext.Provider>
  );
}

export function useData() {
  const context = useContext(DataContext);
  if (context === undefined) {
    throw new Error("useData must be used within a DataProvider");
  }
  return context;
}