
// std
#include <iostream>
// project
#include "bridge.h"
#include "composite.h"
#include "decorator.h"
#include "mediator.h"
#include "strategy.h"
#include "factory_method.h"
#include "visitor.h"
#include "observer.h"
#include "memento.h"
#include "prototype.h"

using namespace std;

int main( int argc, char *argv[] ){

    // 1. Элегантные решения

    // --------------------------
    // Decorator
    // --------------------------
    VisualComponent * vc = new BorderDecorator( new ScrollDecorator(new TextWidget()) );
    vc->draw();

    // --------------------------
    // Visitor
    // --------------------------

    // --------------------------
    // Observer
    // --------------------------

    // --------------------------
    // Mediator
    // --------------------------

    // --------------------------
    // Bridge
    // --------------------------

    // --------------------------
    // State
    // --------------------------


    // 2. Тривиальные решения

    // --------------------------
    // Composite
    // --------------------------

    // --------------------------
    // Strategy
    // --------------------------

    // --------------------------
    // Template method
    // --------------------------

    // --------------------------
    // Singleton TODO
    // --------------------------

    // --------------------------
    // Adapter TODO
    // --------------------------

    // --------------------------
    // Facade TODO
    // --------------------------


    return 0;
}
