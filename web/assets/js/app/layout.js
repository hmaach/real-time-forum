import { getUsers } from "./api.js";
import { handleLogout } from "./auth.js";
import { showDirectMessages } from "./chat.js";
import { showCreatePost } from "./create_post.js";
import { showFeed } from "./feed.js";
import { showErrorPage, formatTime, trimString } from "./utils.js";
import { getOnlineUsers } from "./websocket.js";

export const setupLayout = () => {
    const user = JSON.parse(localStorage.getItem("user"));
    document.body.innerHTML = /*html*/`
        <div id="header-container">
            <header>
                <button id="sidebar-toggle" class="sidebar-toggle">
                    <i class="fa-solid fa-bars"></i>
                </button>
                <div class="title">forum <i class="fa-regular fa-comment-dots"></i></div>
                <div class="header-btns">
                    <div class="new-post-btn">
                        <i class="fa-sharp fa-solid fa-plus"></i>
                        <span>new post</span>
                    </div>
                    <div class="logout-btn">
                        <i class="fa-solid fa-power-off"></i>
                        <span>log out ${user.nickname}</span>
                    </div>
                </div>
            </header>
        </div>

        <div id="body-container">
            <aside id="sidebar" class="sidebar">
                <div class="members-list" id="user-list"></div>
            </aside>

            <main>
            </main>
        </div>
    `;
    // Initialize components
    setupSidebar();
    loadUsers();
    setupEventListeners();
}

const setupSidebar = () => {
    const sidebarToggle = document.getElementById('sidebar-toggle');
    const sidebar = document.getElementById('sidebar');
    const body = document.body;

    // Create and append overlay
    const overlay = document.createElement('div');
    overlay.className = 'sidebar-overlay';
    document.getElementById('body-container').appendChild(overlay);

    const toggleSidebar = () => {
        sidebar.classList.toggle('active');
        overlay.classList.toggle('active');
        body.classList.toggle('sidebar-open');
    };

    sidebarToggle.addEventListener('click', toggleSidebar);
    overlay.addEventListener('click', toggleSidebar);

    // Handle window resize
    window.addEventListener('resize', () => {
        if (window.innerWidth > 900 && sidebar.classList.contains('active')) {
            toggleSidebar();
        }
    });
};

const setupEventListeners = () => {
    const homeBtn = document.querySelector('.title')
    homeBtn.addEventListener('click', showFeed);

    const logoutBtn = document.querySelector('.logout-btn');
    logoutBtn.addEventListener('click', handleLogout);

    const newPostBtn = document.querySelector('.new-post-btn');
    newPostBtn.addEventListener('click', showCreatePost);
};

export const loadUsers = async () => {
    try {
        const token = localStorage.getItem("token");
        const userID = JSON.parse(localStorage.getItem('user')).id
        const response = await getUsers(token);
        const userListContainer = document.querySelector(".members-list");
        userListContainer.innerHTML = "";
        if (response.status !== 200) throw response

        if (!response.users || response.users.length === 0) {
            userListContainer.innerText = "No users to display";
            return;
        }

        response.users.forEach(user => {
            const userElement = document.createElement("div");
            userElement.classList.add("user");
            userElement.setAttribute("data-user-id", user.id);
            var last_message = (user.last_message.sender_id == userID ? "You: " : "") + user.last_message.message
            userElement.innerHTML = /*html*/`
                <div>
                    <img src="https://ui-avatars.com/api/?name=${user.first_name + user.last_name}" alt="profile">
                    <div class="user-status"></div>
                </div>
                <div class="user-list-info">
                    <span>${user.nickname}</span>
                    <div class="last-message">
                        <span>${trimString(last_message, 15)}</span>
                        <span>${formatTime(user.last_message.sent_at)}</span>
                    </div>
                    <div class="typing-indicator-userlist" style="display: none">
                            <span class="typing-text-userlist"></span>
                            <div class="typing-dots">
                                <span>.</span><span>.</span><span>.</span>
                            </div>
                    </div>
                </div>
            `;

            userListContainer.appendChild(userElement);

            userElement.addEventListener('click', () => showDirectMessages(user.id))
        });

        getOnlineUsers()
    } catch (error) {
        console.log(error);
        showErrorPage(error.status, error.response)
    }
}
