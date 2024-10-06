
// project
#include "visitor.h"

// ----------------------------------
// Element
// ----------------------------------
void Floppy::accept( Visitor * _visitor ){
    _visitor->visit( this );
}

void Processor::accept( Visitor * _visitor ){
    _visitor->visit( this );
}

// ----------------------------------
// Visitor
// ----------------------------------
void VisitorPrice::visit( Floppy * _floppy ){
    total += _floppy->price;
}

void VisitorPrice::visit( Processor * _processor ){
    total += _processor->price;
}


void VisitorInventory::visit( Floppy * _floppy ){
    m_namesList.push_back( _floppy->name );
}
void VisitorInventory::visit( Processor * _processor ){
    m_namesList.push_back( _processor->name );
}








