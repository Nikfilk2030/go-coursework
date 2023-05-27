# В этом модуле написан фронтенд как к регистрации, так и к самому чату

## Обязательный список команд

```npm install react-router-dom```

```npm install axios```

```npm install react-bootstrap```

```npm install js-cookie```

```npm install react-router```

> И не забудьте сделать restart после установки библиотек

## Описание файлов
- src: это корневой каталог, где находятся все исходные файлы проекта.

- src/api/Endpoints.js: этот файл определяет все конечные точки API, которые используются в приложении. В этом файле мы определяем базовый URL и специфичные конечные точки для разных функций, таких как вход в систему, регистрация и чат.

- src/components: папка содержит все компоненты React, которые можно использовать в разных частях приложения. Это уменьшает дублирование кода и повышает читаемость кода.

- src/components/Navbar.js: этот компонент отображает верхнюю навигационную панель на всех страницах. Он включает в себя ссылки для навигации между различными страницами.

- src/pages: папка содержит компоненты страниц, каждый из которых соответствует определенному маршруту в приложении.

- src/pages/AboutPage.js: этот файл представляет страницу "О нас". Он содержит информацию о приложении.

- src/pages/ChatPage.js: эта страница отображает чат. Он использует конечную точку API чата для получения и отображения сообщений.

- src/pages/LoginPage.js: эта страница отображает форму входа. Она использует конечную точку API для входа в систему, чтобы проверить учетные данные пользователя и разрешить доступ к чату.

- src/pages/RegisterPage.js: это страница регистрации. Она использует конечную точку API для регистрации для создания нового аккаунта.

---

## Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can't go back!**

If you aren't satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you're on your own.

You don't have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn't feel obligated to use this feature. However we understand that this tool wouldn't be useful if you couldn't customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: [https://facebook.github.io/create-react-app/docs/code-splitting](https://facebook.github.io/create-react-app/docs/code-splitting)

### Analyzing the Bundle Size

This section has moved here: [https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size](https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size)

### Making a Progressive Web App

This section has moved here: [https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app](https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app)

### Advanced Configuration

This section has moved here: [https://facebook.github.io/create-react-app/docs/advanced-configuration](https://facebook.github.io/create-react-app/docs/advanced-configuration)

### Deployment

This section has moved here: [https://facebook.github.io/create-react-app/docs/deployment](https://facebook.github.io/create-react-app/docs/deployment)

### `npm run build` fails to minify

This section has moved here: [https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify](https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify)
