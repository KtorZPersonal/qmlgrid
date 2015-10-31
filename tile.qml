import QtQuick 2.0

Rectangle {
    id: tile
    property int kind: 0

    border { width: 1; color: "#FFFFFF" }

    state: "DEFAULT"
    states: [
        State {
            name: "DEFAULT"
            when: kind == 0
            PropertyChanges { target: tile; color: "#ecf0f1" }
        },

        State {
            name: "BLOCKED"
            when: kind == 1
            PropertyChanges { target: tile; color: "#2c3e50" }
        },

        State {
            name: "VISITED"
            when: kind == 2
            PropertyChanges { target: tile; color: "#7f8c8d" }
        },

        State {
            name: "ACTIVE"
            when: kind == 3
            PropertyChanges { target: tile; color: "#27ae60" }
        },

        State {
            name: "EXTREMITY"
            when: kind == 4
            PropertyChanges { target: tile; color: "#e74c3c" }
        }
    ]
}
