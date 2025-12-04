import './App.css'
import React from "react";
import { UserServiceButtons } from "./components/UserServiceButtons";
import { UserService } from './services/UserService';

function App() {
  const userService = React.useMemo(() => new UserService(), []);


  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-xl shadow-lg text-center space-y-4">
        <h1 className="text-2xl font-bold">User service demo</h1>
        <UserServiceButtons service={userService} />
      </div>
    </div>
  );
}

export default App
