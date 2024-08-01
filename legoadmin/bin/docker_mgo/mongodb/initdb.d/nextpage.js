db.createUser({
    user: "sparkstudio",
    pwd: "maketoo157",
    roles: [
        { role: "readWrite", db: "admin", },
    ],
});