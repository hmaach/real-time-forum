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
            fn.apply(this, arguments);
            last = now;
        }
    };
}

const addcomment = throttle(addcomm, 5000)

function postreaction(postId, reaction) {
    document.getElementById("errorlogin" + postId).innerText = ``
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
                document.getElementById("errorlogin" + postId).innerText = `You must login first!`
                setTimeout(() => {
                    document.getElementById("errorlogin" + postId).innerText = ``
                }, 1000);
            } else if (xhr.status === 400) {
                document.getElementById("errorlogin" + postId).innerText = `Bad request!`
                setTimeout(() => {
                    document.getElementById("errorlogin" + postId).innerText = ``
                }, 1000);
            } else if (xhr.status === 500) {
                document.getElementById("errorlogin" + postId).innerText = `Try again later!`
                setTimeout(() => {
                    document.getElementById("errorlogin" + postId).innerText = ``
                }, 1000);
            }
        };
    }
    xhr.send(`reaction=${reaction}&post_id=${postId}`);
}
function commentreaction(commentid, reaction) {
    document.getElementById("commenterrorlogin" + commentid).innerText = ``
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
                document.getElementById("commenterrorlogin" + commentid).innerText = `You must login first!`
                setTimeout(() => {
                    document.getElementById("commenterrorlogin" + commentid).innerText = ``
                }, 1000);

            } else if (xhr.status === 400) {
                document.getElementById("commenterrorlogin" + commentid).innerText = `bad request!`
                setTimeout(() => {
                    document.getElementById("commenterrorlogin" + commentid).innerText = ``
                }, 1000);
            } else if (xhr.status === 500) {
                document.getElementById("commenterrorlogin" + commentid).innerText = `Try again later!`
                setTimeout(() => {
                    document.getElementById("commenterrorlogin" + commentid).innerText = ``
                }, 1000);
            }
        };
    }
    xhr.send(`reaction=${reaction}&comment_id=${commentid}`);
}


function addcomm(postId) {
    const content = document.getElementById("comment-content");
    const logerror = document.getElementById("errorlogin" + postId)
    if (!content.value) {
        logerror.innerText = 'Please fill in Comment field.';
        setTimeout(() => {
            logerror.innerText = '';
        }, 3000);
        return;
    }

    if (content.value.length > 500) {
        logerror.innerText = 'Comment is too long. Please keep it under 500 characters.';
        setTimeout(() => {
            logerror.innerText = '';
        }, 3000);
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
                document.getElementById("errorlogin" + postId).innerText = `Invalid comment!`
                setTimeout(() => {
                    document.getElementById("errorlogin" + postId).innerText = ``
                }, 1000);
            } else if (xhr.status === 401) {
                document.getElementById("errorlogin" + postId).innerText = `You must login first!`
                setTimeout(() => {
                    document.getElementById("errorlogin" + postId).innerText = ``
                }, 1000);
            } else {
                document.getElementById("errorlogin" + postId).innerText = `Cannot add comment now, try again later!`
                setTimeout(() => {
                    document.getElementById("errorlogin" + postId).innerText = ``
                }, 1000);
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

    if (!title.value || !content.value || categories.childElementCount === 0) {
        logerror.innerText = 'Please fill in all fields and select at least one category.';
        setTimeout(() => {
            logerror.innerText = '';
        }, 3000);
        return;
    }

    if (title.value.length > 100) {
        logerror.innerText = 'Title is too long. Please keep it under 100 characters.';
        setTimeout(() => {
            logerror.innerText = '';
        }, 3000);
        return;
    }

    if (content.value.length > 3000) {
        logerror.innerText = 'Content is too long. Please keep it under 3000 characters.';
        setTimeout(() => {
            logerror.innerText = '';
        }, 3000);
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


                logerror.innerText = 'Post created successfully, redirect to home page in 2s ...'
                logerror.style.color = "green"
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 401) {
                logerror.innerText = 'You are loged out, redirect to login page in 2s...'
                setTimeout(() => {
                    window.location.href = '/login'
                }, 2000)

            } else {
                logerror.innerText = 'Error: check your entries and try again!'
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
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

    const xml = new XMLHttpRequest();
    xml.open("POST", "/signup", true)
    xml.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

    xml.onreadystatechange = function () {
        if (xml.readyState === 4) {
            const logerror = document.querySelector(".errorarea")
            if (xml.status === 200) {
                logerror.innerText = `User ${username.value} created successfully, redirect to login page in 2s ...`
                logerror.style.color = "green"
                setTimeout(() => {
                    window.location.href = '/login'
                }, 2000)

            } else if (xml.status === 302) {
                logerror.innerText = 'You are already loged in, redirect to home page in 2s...'
                logerror.style.color = "green"
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 400) {
                logerror.innerText = 'Error: verify your data and try again!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
            } else if (xml.status === 304) {
                logerror.innerText = 'User already exists!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
            } else {
                logerror.innerText = 'Cannot create user, try again later!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
            }
        }
    }

    // Get form data
    xml.send(`email=${encodeURIComponent(email.value)}&username=${encodeURIComponent(username.value)}&password=${encodeURIComponent(password.value)}&password-confirmation=${encodeURIComponent(passConfirm.value)}`)


}



function login() {
    const username = document.querySelector("#username")
    const password = document.querySelector("#password")

    const xml = new XMLHttpRequest();
    xml.open("POST", "/signin", true)
    xml.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

    xml.onreadystatechange = function () {
        if (xml.readyState === 4) {
            const logerror = document.querySelector(".errorarea")
            if (xml.status === 200) {
                logerror.innerText = `Login in successfully, redirect to home page in 2s ...`
                logerror.style.color = "green"
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 302) {
                logerror.innerText = 'You are already loged in, redirect to home page in 2s...'
                logerror.style.color = "green"
                setTimeout(() => {
                    window.location.href = '/'
                }, 2000)

            } else if (xml.status === 400) {
                logerror.innerText = 'Error: verify your data and try again!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
            } else if (xml.status === 404) {
                logerror.innerText = 'User not found!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
            } else if (xml.status === 401) {
                logerror.innerText = 'Invalid username or password!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
            } else {
                logerror.innerText = 'Cannot log you in now, try again later!'
                logerror.style.color = "red"
                setTimeout(() => {
                    logerror.innerText = ''
                }, 1500)
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

// const formatTime = (timeStr) => {
//     // Parse the input time string
//     const date = new Date(timeStr);

//     return date.toLocaleString('default', {
//         hour: '2-digit',
//         minute: '2-digit',
//         day: '2-digit',
//         month: '2-digit',
//         year: 'numeric',
//     }).replace(',', ' ')
// }

// document.addEventListener("DOMContentLoaded", () => {
//     document.querySelectorAll("[data-timestamp]").forEach((element) => {
//         const time = element.getAttribute("data-timestamp");
//         if (time) {
//             element.textContent = formatTime(time);
//         }
//     });
// });