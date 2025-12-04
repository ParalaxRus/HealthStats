import type { UserService } from "../services/UserService";

interface UserServiceButtonsProps {
  service: UserService;
}

export function UserServiceButtons({ service }: UserServiceButtonsProps) {
  const handleRegisterUser = async () => {
    try {
      const userId = await service.register();
      console.log("UserId:", userId);
    } catch (e) {
      console.error(e);
    }
  };

  const handleListUsers = async () => {
    try {
      await service.list();
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className="space-x-4">
      <button onClick={handleRegisterUser}>Register user</button>
      <button onClick={handleListUsers}>List users</button>
    </div>
  );
}