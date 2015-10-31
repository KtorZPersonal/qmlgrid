import QtQuick 2.0

Rectangle {
    id: root
    color: "#e5e5e5"

    Component.onCompleted: grid.draw(canvas)

    Item {
        id: canvas
        width: parent.width
        anchors { top: parent.top; bottom: parent.bottom }
    }
}
