document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8070/api/v1/user/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password })
        });

        if (!response.ok) {
            throw new Error('登录失败');
        }

        const data = await response.json();
        localStorage.setItem('accessToken', data.accessToken);
        window.location.href = 'api-test.html';
    } catch (error) {
        alert('登录失败：' + error.message);
    }
});