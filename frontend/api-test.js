class RequestComponent {
    constructor(id, config = {}) {
        this.id = id;
        this.config = config;
        this.render();
        this.loadConfig();
    }

    render() {
        const template = `
            <div class="request-component" id="request-${this.id}">
                <div class="component-header">
                    <select class="method-select">
                        <option value="GET">GET</option>
                        <option value="POST">POST</option>
                        <option value="PUT">PUT</option>
                    </select>
                    <button class="delete-button">删除</button>
                </div>
                <input type="text" class="path-input" placeholder="请求路径" />
                <textarea class="json-input" placeholder="JSON数据"></textarea>
                <div class="request-info">
                    <button class="send-button">发送请求</button>
                    <span class="time-info"></span>
                </div>
                <div class="response-container">
                    <textarea class="response-output" readonly placeholder="响应结果将显示在这里"></textarea>
                </div>
            </div>
        `;
        document.getElementById('requestsContainer').insertAdjacentHTML('beforeend', template);

        this.bindEvents();
    }

    deleteComponent() {
        // 从 DOM 中移除组件
        const component = document.getElementById(`request-${this.id}`);
        component.remove();

        // 从 localStorage 中移除配置
        const configs = JSON.parse(localStorage.getItem('apiTestConfigs') || '{}');
        delete configs[this.id];
        localStorage.setItem('apiTestConfigs', JSON.stringify(configs));
    }

    async sendRequest() {
        const component = document.getElementById(`request-${this.id}`);
        const method = component.querySelector('.method-select').value;
        const path = component.querySelector('.path-input').value;
        const jsonData = component.querySelector('.json-input').value;
        const responseOutput = component.querySelector('.response-output');
        const timeInfo = component.querySelector('.time-info');

        // 检查 token
        if (localStorage.getItem('accessToken') === "undefined" || !localStorage.getItem('accessToken')) {
            window.location.href = 'login.html';
            return;
        }

        const startTime = new Date();
        timeInfo.textContent = `开始请求: ${startTime.toLocaleTimeString()}`;

        try {
            let response = await fetch(path, {
                method,
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem('accessToken')
                },
                body: method === 'POST' ? jsonData : undefined,
                redirect: 'follow'
            });

            if (response.url.includes('/api/v1/auth/refresh')) {
                const refreshData = await response.json();
                if (!response.ok) {
                    responseOutput.value = `Status Code: ${response.status}\n\n${JSON.stringify(refreshData, null, 2)}\n\n来自refresh api`;
                    const endTime = new Date();
                    const duration = endTime - startTime;
                    timeInfo.textContent = `完成时间: ${endTime.toLocaleTimeString()} (耗时: ${formatDuration(duration)})`;
                    return;
                }
                localStorage.setItem('accessToken', refreshData.accessToken);

                // 再次发送原始请求
                response = await fetch(path, {
                    method,
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': localStorage.getItem('accessToken')
                    },
                    body: method === 'POST' ? jsonData : undefined
                });
            }

            const endTime = new Date();
            const duration = endTime - startTime;
            const data = await response.json();
            responseOutput.value = `Status Code: ${response.status}\n\n${JSON.stringify(data, null, 2)}`;
            timeInfo.textContent = `完成时间: ${endTime.toLocaleTimeString()} (耗时: ${formatDuration(duration)})`;
        } catch (error) {
            const endTime = new Date();
            const duration = endTime - startTime;
            responseOutput.value = '请求失败：' + error.message;
            timeInfo.textContent = `失败时间: ${endTime.toLocaleTimeString()} (耗时: ${formatDuration(duration)})`;
        }
    }


    bindEvents() {
        const component = document.getElementById(`request-${this.id}`);
        const methodSelect = component.querySelector('.method-select');
        const pathInput = component.querySelector('.path-input');
        const jsonInput = component.querySelector('.json-input');
        const sendButton = component.querySelector('.send-button');
        const deleteButton = component.querySelector('.delete-button');

        // 保存配置的事件
        [methodSelect, pathInput, jsonInput].forEach(element => {
            element.addEventListener('change', () => this.saveConfig());
            element.addEventListener('input', () => this.saveConfig());
        });

        // 发送请求事件
        sendButton.addEventListener('click', () => this.sendRequest());
        // 删除组件事件
        deleteButton.addEventListener('click', () => this.deleteComponent());

    }

    loadConfig() {
        if (this.config) {
            const component = document.getElementById(`request-${this.id}`);
            component.querySelector('.method-select').value = this.config.method || 'GET';
            component.querySelector('.path-input').value = this.config.path || '';
            component.querySelector('.json-input').value = this.config.json || '';
        }
    }

    saveConfig() {
        const component = document.getElementById(`request-${this.id}`);
        const config = {
            method: component.querySelector('.method-select').value,
            path: component.querySelector('.path-input').value,
            json: component.querySelector('.json-input').value
        };

        // 保存到localStorage
        const configs = JSON.parse(localStorage.getItem('apiTestConfigs') || '{}');
        configs[this.id] = config;
        localStorage.setItem('apiTestConfigs', JSON.stringify(configs));
    }
}
// 添加一个格式化持续时间的方法
function formatDuration(duration) {
    if (duration < 1000) {
        return `${duration}ms`;
    } else {
        return `${(duration / 1000).toFixed(2)}s`;
    }
}

// 页面加载时检查登录状态
if (!localStorage.getItem('accessToken')) {
    window.location.href = 'login.html';
}

// 加载保存的配置
const savedConfigs = JSON.parse(localStorage.getItem('apiTestConfigs') || '{}');
if (Object.keys(savedConfigs).length > 0) {
    // 如果有保存的配置，则加载它们
    Object.entries(savedConfigs).forEach(([id, config]) => {
        new RequestComponent(id, config);
    });
} else {
    // 如果没有保存的配置，创建一个新的请求组件
    new RequestComponent(Date.now().toString());
}

// 添加新请求按钮事件
document.getElementById('addRequest').addEventListener('click', () => {
    const id = Date.now().toString();
    new RequestComponent(id);
});

// 布局切换按钮
const singleColumnBtn = document.getElementById('singleColumn');
const doubleColumnBtn = document.getElementById('doubleColumn');
const requestsContainer = document.getElementById('requestsContainer');

// 保存布局偏好到 localStorage
function saveLayoutPreference(layout) {
    localStorage.setItem('layoutPreference', layout);
}

// 加载布局偏好
function loadLayoutPreference() {
    const layout = localStorage.getItem('layoutPreference') || 'single-column';
    requestsContainer.className = layout;
    if (layout === 'single-column') {
        singleColumnBtn.classList.add('active');
        doubleColumnBtn.classList.remove('active');
    } else {
        doubleColumnBtn.classList.add('active');
        singleColumnBtn.classList.remove('active');
    }
}

// 绑定布局切换事件
singleColumnBtn.addEventListener('click', () => {
    requestsContainer.className = 'single-column';
    singleColumnBtn.classList.add('active');
    doubleColumnBtn.classList.remove('active');
    saveLayoutPreference('single-column');
});

doubleColumnBtn.addEventListener('click', () => {
    requestsContainer.className = 'double-column';
    doubleColumnBtn.classList.add('active');
    singleColumnBtn.classList.remove('active');
    saveLayoutPreference('double-column');
});

// 页面加载时应用保存的布局偏好
loadLayoutPreference();