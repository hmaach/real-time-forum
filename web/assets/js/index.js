window.addEventListener('resize', () => {
    if (document.body.clientWidth > 600) {
        document.querySelector('.mobile-nav').style.display = 'none';
    }
})


function throttle(fn, delay) {
    let last = 0;
    return function () {
        const now = +new Date();
        if (now - last > delay) {
            fn.apply(null, arguments);
            last = now;
        }
    };
}

const addcomment = throttle(addcomm, 5000)

function postreaction(postId, reaction) {
    const logerror = document.getElementById("errorlogin" + postId)
    logerror.innerText = ``
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/postreaction", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);
                document.getElementById("likescount" + postId).innerHTML = `<i
                    class="fa-regular fa-thumbs-up"></i>${response.likesCount}`;
                document.getElementById("dislikescount" + postId).innerHTML = `<i
                    class="fa-regular fa-thumbs-down"></i>${response.dislikesCount}`;
            } else if (xhr.status === 401) {
                writeError(logerror,"red",`You must login first!`,1000)
            } else if (xhr.status === 400) {
                writeError(logerror,"red",`Bad request!`,1000)
            } else if (xhr.status === 500) {
                writeError(logerror,"red",`Try again later!`,1000)
            }
        };
    }
    xhr.send(`reaction=${reaction}&post_id=${postId}`);
}
function commentreaction(commentid, reaction) {
    const logerror = document.getElementById("commenterrorlogin" + commentid)
    logerror.innerText = ``
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/commentreaction", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);
                document.getElementById("commentlikescount" + commentid).innerHTML = `<i
                    class="fa-regular fa-thumbs-up"></i>${response.commentlikesCount}`;
                document.getElementById("commentdislikescount" + commentid).innerHTML = `<i
                    class="fa-regular fa-thumbs-down"></i>${response.commentdislikesCount}`;
            } else if (xhr.status === 401) {
                writeError(logerror,"red",`You must login first!`,1000)
                
            } else if (xhr.status === 400) {
                writeError(logerror,"red",`bad request!`,1000)
            
            } else if (xhr.status === 500) {
                writeError(logerror,"red",`Try again later!`,1000)
            }
        };
    }
    xhr.send(`reaction=${reaction}&comment_id=${commentid}`);
}


function addcomm(postId) {
    const content = document.getElementById("comment-content");
    const logerror = document.getElementById("errorlogin" + postId)
    if (!content.value) {
        writeError(logerror,"red",'Please fill in Comment field.',3000)
        return;
    }

    if (content.value.length > 500) {
        writeError(logerror,"red",'Comment is too long. Please keep it under 500 characters.',3000)
        return;
    }

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/addcommentREQ", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);
                const comment = document.createElement("div")
                comment.innerHTML = `
                 <div class="comment">
            <div class="comment-header">
                <p class="comment-user">`+ response.username + `</p>
                <span></span>
                <p class="comment-time">`+ response.created_at + ` </p>
            </div>
            <div class="comment-body">
                <p class="comment-content">`+ response.content + ` </p>
            </div>
            <div class="comment-footer">
                <button id="commentlikescount`+ response.ID + `" onclick="commentreaction('` + response.ID + `','like')"
                    class="comment-like"><i class="fa-regular fa-thumbs-up"></i>`+ response.likes + `</button>
                <button id="commentdislikescount`+ response.ID + `" onclick="commentreaction('` + response.ID + `','dislike')"
                    class="comment-dislike"><i class="fa-regular fa-thumbs-down"></i>`+ response.dislikes + `</button>
            </div>
            <span style="color:red" id="commenterrorlogin`+ response.ID + `"></span>
        </div>
                `
                document.getElementsByClassName("comments")[0].prepend(comment)
                document.getElementsByClassName("post-comments")[0].innerHTML = `<i class="fa-regular fa-comment"></i>` + response.commentscount
                content.value = ""
            } else if (xhr.status === 400) {
                writeError(logerror,"red",`Invalid comment!`,1000)
            } else if (xhr.status === 401) {
                writeError(logerror,"red",`You must login first!`,1000)
            } else {
                writeError(logerror,"red",`Cannot add comment now, try again later!`,1000)

            }
        };
    }
    xhr.send(`postid=${postId}&comment=${encodeURIComponent(content.value)}`);
}

const select = document.getElementById('categories-select');
if (select) {

    select.addEventListener('change', (e) => {
        // Parse the value as JSON to extract id and label
        const selectedValue = JSON.parse(e.target.value);
        const { id, label } = selectedValue;

        // create the elemenet for the category
        const span = document.createElement('span');
        span.textContent = label;
        span.classList.add('selected-category');

        // Add a remove button to the span
        const removeBtn = document.createElement('span');
        removeBtn.textContent = '×';
        removeBtn.classList.add('remove-category');
        removeBtn.addEventListener('click', () => {
            span.remove();
            input.remove();
            // Re-enable the corresponding option in the select
            Array.from(e.target.options).find(option => {
                try {
                    const optionValue = JSON.parse(option.value);
                    return optionValue.id === id;
                } catch {
                    return false;
                }
            }).disabled = false;
        });

        span.appendChild(removeBtn);

        // create hidden input to hold the id of selected category
        const input = document.createElement('input')
        input.type = 'hidden';
        input.value = id
        input.name = 'categories'

        // add the elements (span and hidden input) 
        // at the first position of the categories container
        const categoriesContainer = document.querySelector('.selected-categories');
        categoriesContainer.append(input, span);

        // disable the option selected in the select
        e.target.options[e.target.selectedIndex].disabled = true;

        // Reset the select 
        e.target.selectedIndex = 0;
    });
}

async function pagination(dir, data) {
    const path = window.location.pathname
    if (dir === "next" && data) {
        const page = +document.querySelector(".currentpage").innerText + 1
        window.location.href = path + "?PageID=" + page;
    }

    if (dir === "back" && document.querySelector(".currentpage").innerText > "1") {
        const page = +document.querySelector(".currentpage").innerText - 1
        window.location.href = path + "?PageID=" + page;
    }
}



function CreatPost() {
    const title = document.querySelector(".create-post-title")
    const content = document.querySelector(".content")
    const categories = document.querySelector(".selected-categories")
    const logerror = document.querySelector(".errorarea")

    if (title.value.trim() == "" || content.value.trim() == "" || categories.childElementCount === 0) {
        writeError(logerror,"red",'No empty entries allowed!',3000)
        return;
    }

    if (title.value.length > 100) {
        writeError(logerror,"red",'Title is too long. Please keep it under 100 characters.',3000)
        return;
    }

    if (content.value.length > 3000) {
        writeError(logerror,"red",'Content is too long. Please keep it under 3000 characters.',3000)
        return;
    }


    let cateris = new Array()
    Array.from(categories.getElementsByTagName('input')).forEach((x) => {
        cateris.push(x.value)
    })
    const xml = new XMLHttpRequest();
    xml.open("POST", "/post/createpost", true)
    xml.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

    xml.onreadystatechange = function () {
        if (xml.readyState === 4) {
            if (xml.status === 200) {
                const btn = document.getElementById("create-post-btn")
                document.getElementById("publish-post-icon").style.display = "none"
                document.getElementById("publish-post-circle").style.display = "inline-block"
                btn.disabled = true
                btn.style.background = "grey"
                btn.style.cursor = "not-allowed"


                writeError(logerror,"green",'Post created successfully, redirect to home page in 2s ...',2000)
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 401) {
                writeError(logerror,"red",'You are loged out, redirect to login page in 2s...',2000)
                setTimeout(() => {
                    window.location.href = '/login'
                }, 2000)

            } else if (xml.status === 400) {
                writeError(logerror,"red",'Bad request!',1500)
            } else {
                writeError(logerror,"red",'Error: check your entries and try again!',1500)
            }
        }
    }

    // Get form data
    xml.send(`title=${encodeURIComponent(title.value)}&content=${encodeURIComponent(content.value)}&categories=${cateris}`)
}


function register() {
    const email = document.querySelector("#email")
    const username = document.querySelector("#username")
    const password = document.querySelector("#password")
    const passConfirm = document.querySelector("#password-confirmation")
    const logerror = document.querySelector(".errorarea")

    if (username.value.length < 4 || username.value.includes(" ")){
        writeError(logerror,"red","Username too short! or have space",1500)
        return
    }

    if (password.value.length < 6){
        writeError(logerror,"red","password too short!",1500)
        return
    }

    if (password.value != passConfirm.value){
        writeError(logerror,"red","password and password confirmation are not identical",1500)
        return
    }


    const xml = new XMLHttpRequest();
    xml.open("POST", "/signup", true)
    xml.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

    xml.onreadystatechange = function () {
        if (xml.readyState === 4) {
            if (xml.status === 200) {
                writeError(logerror,"green",`User ${username.value} created successfully, redirect to login page in 2s ...`,2000)
                setTimeout(() => {
                    window.location.href = '/login'
                }, 2000)

            } else if (xml.status === 302) {
                writeError(logerror,"green",'You are already loged in, redirect to home page in 2s...',2000)
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 400) {
                writeError(logerror,"red",'Error: verify your data and try again!',1500)
    
            } else if (xml.status === 304) {
                writeError(logerror,"red",'User already exists!',1500)
    
            } else {
                writeError(logerror,"red",'Cannot create user, try again later!',1500)
            }
        }
    }

    // Get form data
    xml.send(`email=${encodeURIComponent(email.value)}&username=${encodeURIComponent(username.value)}&password=${encodeURIComponent(password.value)}&password-confirmation=${encodeURIComponent(passConfirm.value)}`)


}



function login() {
    const username = document.querySelector("#username")
    const password = document.querySelector("#password")
    const logerror = document.querySelector(".errorarea")

    if (username.value.length < 4) {
        writeError(logerror,"red","Username too short!",1500)
        return
    }
    if (password.value.length < 6) {
        writeError(logerror,"red","Password too short!",1500)
        return
    }


    const xml = new XMLHttpRequest();
    xml.open("POST", "/signin", true)
    xml.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

    xml.onreadystatechange = function () {
        if (xml.readyState === 4) {
            if (xml.status === 200) {
                writeError(logerror,"green",`Login in successfully, redirect to home page in 2s ...`,2000)
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)
            } else if (xml.status === 302) {
                writeError(logerror,"green",'You are already loged in, redirect to home page in 2s...',2000)
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 400) {
                writeError(logerror,"red",'Error: verify your data and try again!',1500)
            } else if (xml.status === 404) {
                writeError(logerror,"red",'User not found!',1500)
            } else if (xml.status === 401) {
                writeError(logerror,"red",'Invalid username or password!',1500)
            } else {
                writeError(logerror,"red",'Cannot log you in now, try again later!',1500)
            }
        }
    }

    // Get form data
    xml.send(`username=${encodeURIComponent(username.value)}&password=${encodeURIComponent(password.value)}`)
}

const displayMobileNav = (e) => {
    const nav = document.querySelector('.mobile-nav')
    nav.style.display = 'block'
}

const closeMobileNav = (e) => {
    const nav = document.querySelector('.mobile-nav')
    nav.style.display = 'none'
}

function writeError(targetDiv,color,errormsg,delay) {
    targetDiv.innerText = errormsg
    targetDiv.style.color = color
    setTimeout(() => {
        targetDiv.innerText = ''
    }, delay)
}